package bun

import (
	"context"
	"time"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/models"
	"github.com/uptrace/bun"
)

func (c *Client) CreateAccount(ctx context.Context, account *models.Account) db.Error {
	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now

	query := c.db.NewInsert().
		Model(&account)

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}

func (c *Client) ReadAccountByToken(ctx context.Context, token string) (*models.Account, db.Error) {
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

func (c *Client) ReadAccountByUserID(ctx context.Context, userID string) (*models.Account, db.Error) {
	account := new(models.Account)
	query := newAccountQ(c.db, account).
		Where("account.user_id = ?", userID)

	if err := query.Scan(ctx); err != nil {
		return nil, processError(err)
	}

	return account, nil
}

func (c *Client) UpdateAccount(ctx context.Context, account *models.Account) db.Error {
	account.UpdatedAt = time.Now()

	query := c.db.NewUpdate().
		Model(&account).
		WherePK()

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}

func newAccountQ(c bun.IDB, account *models.Account) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(account)
}
