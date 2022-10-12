package models

import "time"

type LocalThreepidAssociation struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	Medium     string    `validate:"-" bun:",nullzero,notnull"`
	Address    string    `validate:"-" bun:",nullzero,notnull"`
	MXID       string    `validate:"-" bun:",nullzero,notnull"`
	LookupHash string    `validate:"-" bun:",nullzero,notnull"`
	NotBefore  time.Time `validate:"-" bun:",nullzero,notnull"`
	NotAfter   time.Time `validate:"-" bun:",nullzero,notnull"`
}
