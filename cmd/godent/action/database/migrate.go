package database

import (
	"context"
	"github.com/tyrm/godent/cmd/godent/action"
	"github.com/tyrm/godent/internal/db/bun"
)

// Migrate runs database migrations
var Migrate action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Migrate")

	// create database client
	l.Info("running database migration")
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

	err = dbClient.DoMigration(ctx)
	if err != nil {
		l.Errorf("migration: %s", err.Error())

		return err
	}

	return nil
}
