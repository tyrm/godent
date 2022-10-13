package models

import "time"

type Token struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`

	Token     string   `bun:",nullzero"`
	AccountID int64    `bun:",nullzero,notnull"`
	Account   *Account `bun:"rel:belongs-to,join:account_id=id"`
}
