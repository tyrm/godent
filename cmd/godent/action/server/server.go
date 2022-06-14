package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/cmd/godent/action"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db/bun"
	"github.com/tyrm/godent/internal/db/memcache"
	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/http/versions"
	"github.com/tyrm/godent/internal/kv/redis"
	"github.com/tyrm/godent/internal/util"
)

// Start starts the server.
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")

	l.Infof("starting")
	l.Infof("creating db client")
	dbClient, err := bun.New(ctx)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	l.Infof("creating db cache client")
	cachedDBClient, err := memcache.New(ctx, dbClient)
	if err != nil {
		l.Errorf("db-cachemem: %s", err.Error())

		return err
	}
	defer func() {
		err := cachedDBClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	redisClient, err := redis.New(ctx)
	if err != nil {
		l.Errorf("redis: %s", err.Error())

		return err
	}
	defer func() {
		err := redisClient.Close(ctx)
		if err != nil {
			l.Errorf("closing redis: %s", err.Error())
		}
	}()

	// create http server
	l.Debug("creating http server")
	httpServer, err := gdhttp.NewServer(ctx)
	if err != nil {
		l.Errorf("http httpServer: %s", err.Error())

		return err
	}

	// create web modules
	var webModules []gdhttp.Module
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleVersions) {
		l.Infof("adding wellknown module")
		webMod, err := versions.New(ctx)
		if err != nil {
			l.Errorf("wellknown module: %s", err.Error())

			return err
		}
		webModules = append(webModules, webMod)
	}

	// add modules to server
	for _, mod := range webModules {
		mod.SetServer(httpServer)
		err := mod.Route(httpServer)
		if err != nil {
			l.Errorf("loading %s module: %s", mod.Name(), err.Error())

			return err
		}
	}

	// ** start application **
	errChan := make(chan error)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	stopSigChan := make(chan os.Signal, 1)
	signal.Notify(stopSigChan, syscall.SIGINT, syscall.SIGTERM)

	// start webserver
	go func(s *gdhttp.Server, errChan chan error) {
		l.Debug("starting http server")
		err := s.Start()
		if err != nil {
			errChan <- fmt.Errorf("http server: %s", err.Error())
		}
	}(httpServer, errChan)

	// wait for event
	select {
	case sig := <-stopSigChan:
		l.Infof("got sig: %s", sig)
	case err := <-errChan:
		l.Fatal(err.Error())
	}

	l.Infof("done")

	return nil
}
