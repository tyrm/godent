package bun

import (
	"context"
	"time"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/models"
)

func (c *Client) CreateToken(ctx context.Context, token *models.Token) db.Error {
	now := time.Now()
	token.CreatedAt = now

	query := c.db.NewInsert().
		Model(&token)

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}

func (c *Client) DeleteToken(ctx context.Context, token *models.Token) db.Error {
	query := c.db.NewDelete().
		Model(&token).
		WherePK()

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}
