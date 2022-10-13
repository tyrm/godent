package models

import "time"

type Account struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	UserID         string `bun:",nullzero,notnull"`
	ConsentVersion string `bun:",nullzero"`
}
