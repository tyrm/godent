package models

import "time"

type GlobalThreepidAssociation struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	Medium       string    `bun:",nullzero,notnull"`
	Address      string    `bun:",nullzero,notnull"`
	MXID         string    `bun:",nullzero,notnull"`
	LookupHash   string    `bun:",nullzero,notnull"`
	NotBefore    time.Time `bun:",nullzero,notnull"`
	NotAfter     time.Time `bun:",nullzero,notnull"`
	OriginServer string    `bun:",nullzero,notnull"`
	OriginID     int64     `bun:",nullzero,notnull"`
	SGAssoc      string    `bun:",nullzero,notnull"`
}
