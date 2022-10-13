package models

import "time"

type InviteToken struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	Medium   string    `bun:",nullzero,notnull"`
	Address  string    `bun:",nullzero,notnull"`
	RoomID   string    `bun:",nullzero,notnull"`
	Sender   string    `bun:",nullzero,notnull"`
	Token    string    `bun:",nullzero,notnull"`
	Received time.Time `bun:",nullzero"`
	Sent     time.Time `bun:",nullzero"`
}
