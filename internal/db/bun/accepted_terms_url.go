package bun

import (
	"context"
	"time"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/models"
	"github.com/uptrace/bun"
)

func (c *Client) CreateAcceptedTermsURL(ctx context.Context, acceptedTermsURLs ...*models.AcceptedTermsURL) db.Error {
	now := time.Now()
	setAcceptedTermsURLsCreatedAt(&acceptedTermsURLs, now)
	setAcceptedTermsURLsUpdatedAt(&acceptedTermsURLs, now)

	query := c.db.NewInsert().
		Model(&acceptedTermsURLs).
		On("CONFLICT DO NOTHING")

	if _, err := query.Exec(ctx); err != nil {
		return processError(err)
	}

	return nil
}

func (c *Client) ReadAcceptedTermsURLForAccount(ctx context.Context, accountID int64) ([]*models.AcceptedTermsURL, db.Error) {
	var acceptedTermsURLs []*models.AcceptedTermsURL
	query := newAcceptedTermsURLsQ(c.db, &acceptedTermsURLs).
		Where("accepted_term_url.account_id = ?", accountID)

	if err := query.Scan(ctx); err != nil {
		return nil, processError(err)
	}

	return acceptedTermsURLs, nil
}

func newAcceptedTermsURLsQ(c bun.IDB, acceptedTermsURLs *[]*models.AcceptedTermsURL) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(acceptedTermsURLs)
}

func setAcceptedTermsURLsCreatedAt(acceptedTermsURLs *[]*models.AcceptedTermsURL, t time.Time) {
	for _, a := range *acceptedTermsURLs {
		a.CreatedAt = t
	}
}

func setAcceptedTermsURLsUpdatedAt(acceptedTermsURLs *[]*models.AcceptedTermsURL, t time.Time) {
	for _, a := range *acceptedTermsURLs {
		a.UpdatedAt = t
	}
}
