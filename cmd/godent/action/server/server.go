package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/godent/internal/cache"
	"github.com/tyrm/godent/internal/cache/memory"
	"github.com/tyrm/godent/internal/http/pubkey"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/cmd/godent/action"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db/bun"
	"github.com/tyrm/godent/internal/fc"
	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/http/account"
	"github.com/tyrm/godent/internal/http/status"
	"github.com/tyrm/godent/internal/http/terms"
	"github.com/tyrm/godent/internal/http/versions"
	"github.com/tyrm/godent/internal/kv/redis"
	logic "github.com/tyrm/godent/internal/logic/v1"
	"github.com/uptrace/uptrace-go/uptrace"
)

// Start starts the server.
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")

	if logrus.GetLevel() == logrus.TraceLevel {
		logConfig(l)
	}

	l.Infof("starting")

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(viper.GetString(config.Keys.ApplicationName)),
		uptrace.WithServiceVersion(viper.GetString(config.Keys.SoftwareVersion)),
	)

	l.Info("creating db client")
	dbClient, err := bun.New(ctx)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	l.Info("creating redis client")
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

	// cache
	var cacher cache.Cache
	switch strings.ToLower(viper.GetString(config.Keys.CacheStore)) {
	case config.CacheStoreMemory:
		l.Info("creating memory cache")
		memoryCache, err := memory.New(ctx)
		if err != nil {
			l.Errorf("memory cacher: %s", err.Error())

			return err
		}

		cacher = memoryCache
	default:
		l.Warnf("unknowm cacher type '%s', defaulting to memory cache", viper.GetString(config.Keys.CacheStore))
		memoryCache, err := memory.New(ctx)
		if err != nil {
			l.Errorf("memory cacher: %s", err.Error())

			return err
		}

		cacher = memoryCache

	}

	// http clients
	httpRetryClient := gdhttp.NewRetryClient()
	federatingClient := fc.New(cacher, httpRetryClient)

	// logic
	logicMod := logic.New(
		dbClient,
		federatingClient,
	)

	// create http server
	l.Debug("creating http server")
	httpServer, err := gdhttp.NewServer(ctx)
	if err != nil {
		l.Errorf("http httpServer: %s", err.Error())

		return err
	}

	// create web modules
	var webModules []gdhttp.Module

	l.Info("adding accound module")
	httpAccount, err := account.New(ctx, federatingClient, logicMod)
	if err != nil {
		l.Errorf("account module: %s", err.Error())

		return err
	}
	webModules = append(webModules, httpAccount)

	l.Info("adding status module")
	httpStatus, err := status.New(ctx)
	if err != nil {
		l.Errorf("status module: %s", err.Error())

		return err
	}
	webModules = append(webModules, httpStatus)

	l.Info("adding pubkey module")
	httpPubkey, err := pubkey.New(ctx, logicMod)
	if err != nil {
		l.Errorf("pubkey module: %s", err.Error())

		return err
	}
	webModules = append(webModules, httpPubkey)

	l.Info("adding terms module")
	httpTerms, err := terms.New(ctx, logicMod)
	if err != nil {
		l.Errorf("terms module: %s", err.Error())

		return err
	}
	webModules = append(webModules, httpTerms)

	l.Info("adding versions module")
	httpVersions, err := versions.New(ctx)
	if err != nil {
		l.Errorf("versions module: %s", err.Error())

		return err
	}
	webModules = append(webModules, httpVersions)

	// add modules to server
	for _, mod := range webModules {
		err := mod.Route(httpServer)
		if err != nil {
			l.Errorf("loading %s module: %s", mod.Name(), err.Error())

			return err
		}
	}

	// start logic module
	err = logicMod.Start(ctx)
	if err != nil {
		l.Errorf("logic module: %s", err.Error())

		return err
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
