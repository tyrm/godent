package models

import "time"

type EphemeralPublicKeys struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	PublicKey   string    `validate:"-" bun:",nullzero,notnull"`
	VerifyCount int64     `validate:"-" bun:",notnull"`
	Persistence time.Time `validate:"-" bun:",nullzero"`
}
