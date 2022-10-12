package bun

import (
	"context"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/db/bun/migrations"
	"github.com/uptrace/bun/migrate"
)

// Close closes the bun db connection.
func (c *Client) Close(_ context.Context) db.Error {
	l := logger.WithField("func", "Close")
	l.Info("closing db connection")

	return c.db.Close()
}

// DoMigration runs schema migrations on the database.
func (c *Client) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		if err.Error() == "migrate: there are no any migrations" {
			return nil
		}

		return err
	}

	if group.ID == 0 {
		l.Info("there are no new migrations to run")

		return nil
	}

	l.Infof("migrated database to %s", group)

	return nil
}
