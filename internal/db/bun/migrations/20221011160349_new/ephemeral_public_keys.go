package models

import "time"

type EphemeralPublicKey struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	PublicKey   string    `bun:",nullzero,notnull"`
	VerifyCount int64     `bun:",notnull"`
	Persistence time.Time `bun:",nullzero"`
}
