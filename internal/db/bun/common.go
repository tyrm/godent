package bun

import (
	"context"
	"errors"
	"fmt"
	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/db/bun/migrations"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/migrate"
)

// Close closes the bun db connection.
func (c *Client) Close(_ context.Context) db.Error {
	l := logger.WithField("func", "Close")
	l.Info("closing db connection")

	return c.bun.Close()
}

// Create inserts an object into the database.
func (c *Client) Create(ctx context.Context, i interface{}) db.Error {
	_, err := c.bun.NewInsert().Model(i).Exec(ctx)
	if err != nil {
		logger.WithField("func", "Create").Errorf("db: %s", err.Error())
	}

	return c.bun.ProcessError(err)
}

// DoMigration runs schema migrations on the database.
func (c *Client) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.bun.DB, migrations.Migrations)

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

// LoadTestData adds test data to the database.
func (c *Client) LoadTestData(ctx context.Context) db.Error {
	l := logger.WithField("func", "LoadTestData")
	l.Debugf("adding test data")

	// Truncate
	modelList := []interface{}{}

	for _, m := range modelList {
		l.Debugf("truncating %T", m)
		_, err := c.bun.NewTruncateTable().Model(m).Exec(ctx)
		if err != nil {
			l.Errorf("truncating %T: %s", m, err.Error())

			return err
		}
	}

	// fix sequences
	sequences := []struct {
		table        string
		currentValue int
	}{}

	switch c.bun.Dialect().Name() {
	case dialect.SQLite:
		// nothing to do
	case dialect.PG:
		for _, s := range sequences {
			_, err := c.bun.Exec("SELECT setval(?, ?, true);", fmt.Sprintf("%s_id_seq", s.table), s.currentValue)
			if err != nil {
				l.Errorf("can't update sequence for %s: %s", s.table, err.Error())

				return err
			}
		}
	default:
		return errors.New("unknown dialect")
	}

	return nil
}

// ReadByID returns a model by its ID.
func (c *Client) ReadByID(ctx context.Context, id int64, i interface{}) db.Error {
	q := c.bun.NewSelect().Model(i).Where("id = ?", id)
	err := q.Scan(ctx)

	return c.bun.ProcessError(err)
}

// ResetCache does nothing. This module contains no cache.
func (*Client) ResetCache(_ context.Context) db.Error {
	return nil
}

// Update updates stored data.
func (c *Client) Update(ctx context.Context, i interface{}) db.Error {
	q := c.bun.NewUpdate().Model(i).WherePK()
	_, err := q.Exec(ctx)

	return c.bun.ProcessError(err)
}
