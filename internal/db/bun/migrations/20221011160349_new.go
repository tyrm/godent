package migrations

import (
	"context"

	models "github.com/tyrm/godent/internal/db/bun/migrations/20221011160349_new"

	"github.com/uptrace/bun"
)

func init() {
	l := logger.WithField("migration", "20221011160349")

	modelList := []interface{}{
		&models.EphemeralPublicKey{},
	}

	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			for _, i := range modelList {
				query := tx.NewCreateTable().Model(i).IfNotExists()

				l.Infof(CreatingTable, query.GetTableName())
				if _, err := query.Exec(ctx); err != nil {
					l.Errorf(CreatingTableErr, query.GetTableName(), err.Error())

					return err
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			for _, i := range modelList {
				query := tx.NewDropTable().Model(i)

				l.Infof(DroppingTable, query.GetTableName())
				if _, err := query.Exec(ctx); err != nil {
					l.Errorf(DroppingTableErr, query.GetTableName(), err.Error())

					return err
				}
			}

			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
