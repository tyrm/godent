package models

import "time"

type AcceptedTermsURLs struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	AccountID int64    `validate:"-" bun:",nullzero,notnull"`
	Account   *Account `validate:"-" bun:"rel:belongs-to,join:account_id=id"`
	URL       string   `validate:"-" bun:",nullzero,notnull"`
}
