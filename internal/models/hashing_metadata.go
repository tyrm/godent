package models

import "time"

type HashingMetadata struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	LookupPepper string `bun:",nullzero"`
}
