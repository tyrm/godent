package models

import "time"

type PeerKey struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull"`

	PeerID    int64  `validate:"-" bun:",nullzero,notnull"`
	Peer      *Peer  `validate:"-" bun:"rel:belongs-to,join:peer_id=id"`
	Algo      string `validate:"-" bun:",nullzero,notnull"`
	PublicKey string `validate:"-" bun:",nullzero,notnull"`
}
