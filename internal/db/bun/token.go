package bun

import (
	"context"
	"time"

	"github.com/uptrace/bun"

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

func (c *Client) ReadTokenByToken(ctx context.Context, t string) (*models.Token, db.Error) {
	token := new(models.Token)
	query := newTokenQ(c.db, token).
		Where("token.token = ?", t)

	if err := query.Scan(ctx); err != nil {
		return nil, processError(err)
	}

	return token, nil
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

func newTokenQ(c bun.IDB, token *models.Token) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(token).
		Relation("Account")
}
