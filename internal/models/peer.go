package models

import "time"

type Peer struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	Name                string `validate:"-" bun:",nullzero,notnull"`
	Port                int    `validate:"-" bun:",nullzero"`
	LastSentVersion     int    `validate:"-" bun:",nullzero"`
	LastPokeSucceededAt int    `validate:"-" bun:",nullzero"`
	Active              bool   `validate:"-" bun:",notnull"`
}
