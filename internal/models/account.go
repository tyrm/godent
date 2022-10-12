package models

import "time"

type Account struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	UserID         string `validate:"-" bun:",nullzero,notnull"`
	ConsentVersion string `validate:"-" bun:",nullzero"`
}
