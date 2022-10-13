package models

import "time"

type ThreePIDTokenAuth struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	ValidationSession int64  `bun:",nullzero,notnull"`
	Token             string `bun:",nullzero,notnull"`
	SendAttemptNumber int64  `bun:",nullzero,notnull"`
}
