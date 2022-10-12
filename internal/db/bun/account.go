package bun

import (
	"context"

	"github.com/tyrm/godent/internal/models"
	"github.com/uptrace/bun"
)

func (c *Client) ReadAccountByToken(ctx context.Context, token string) (*models.Account, error) {
	account := new(models.Account)
	query := newAccountQ(c.db, account).
		Join("JOIN tokens").
		JoinOn("account.id = tokens.account_id").
		Where("tokens.token = ?", token)

	if err := query.Scan(ctx); err != nil {
		return nil, processError(err)
	}

	return account, nil
}

func newAccountQ(c bun.IDB, account *models.Account) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(account).
		Relation("Instance").
		Relation("Groups")
}
