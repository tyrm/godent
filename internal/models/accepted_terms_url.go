package models

import "time"

type AcceptedTermsURL struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	AccountID int64    `bun:",nullzero,notnull"`
	Account   *Account `bun:"rel:belongs-to,join:account_id=id"`
	URL       string   `bun:",nullzero,notnull"`
}
