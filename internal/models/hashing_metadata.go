package models

import "time"

type HashingMetadata struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	LookupPepper string `validate:"-" bun:",nullzero"`
}
