package models

import "time"

type Peer struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	Name                string `bun:",nullzero,notnull"`
	Port                int    `bun:",nullzero"`
	LastSentVersion     int    `bun:",nullzero"`
	LastPokeSucceededAt int    `bun:",nullzero"`
	Active              bool   `bun:",notnull"`
}
