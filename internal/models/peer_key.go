package models

import "time"

type PeerKey struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull"`
	UpdatedAt time.Time `bun:",nullzero,notnull"`

	PeerID    int64  `bun:",nullzero,notnull"`
	Peer      *Peer  `bun:"rel:belongs-to,join:peer_id=id"`
	Algo      string `bun:",nullzero,notnull"`
	PublicKey string `bun:",nullzero,notnull"`
}
