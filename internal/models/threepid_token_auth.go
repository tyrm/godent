package models

import "time"

type ThreePIDTokenAuth struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	ValidationSession int64  `validate:"-" bun:",nullzero,notnull"`
	Token             string `validate:"-" bun:",nullzero,notnull"`
	SendAttemptNumber int64  `validate:"-" bun:",nullzero,notnull"`
}
