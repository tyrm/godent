package models

import "time"

type InviteToken struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	Medium   string    `validate:"-" bun:",nullzero,notnull"`
	Address  string    `validate:"-" bun:",nullzero,notnull"`
	RoomID   string    `validate:"-" bun:",nullzero,notnull"`
	Sender   string    `validate:"-" bun:",nullzero,notnull"`
	Token    string    `validate:"-" bun:",nullzero,notnull"`
	Received time.Time `validate:"-" bun:",nullzero"`
	Sent     time.Time `validate:"-" bun:",nullzero"`
}
