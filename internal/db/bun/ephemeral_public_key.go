package bun

import (
	"context"
	"time"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/models"
)

func (c *Client) CreateEphemeralPublicKey(ctx context.Context, ephemeralPublicKey *models.EphemeralPublicKey) db.Error {
	now := time.Now()
	ephemeralPublicKey.CreatedAt = now
	ephemeralPublicKey.UpdatedAt = now

	query := c.db.NewInsert().
		Model(&ephemeralPublicKey)

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}

func (c *Client) IncEphemeralPublicKeyVerifyCountByPublicKey(ctx context.Context, publicKey string) (int64, db.Error) {
	query := c.db.NewUpdate().
		Model((*models.EphemeralPublicKey)(nil)).
		Set("verify_count = verify_count + 1").
		Where("public_key = ?", publicKey)

	result, err := query.Exec(ctx)
	if err != nil {
		return 0, processError(err)
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return 0, processError(err)
	}

	return rowsEffected, nil
}
