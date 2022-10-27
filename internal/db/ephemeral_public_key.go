package db

import (
	"context"

	"github.com/tyrm/godent/internal/models"
)

type EphemeralPublicKey interface {
	CreateEphemeralPublicKey(ctx context.Context, ephemeralPublicKey *models.EphemeralPublicKey) Error
	IncEphemeralPublicKeyVerifyCountByPublicKey(ctx context.Context, publicKey string) (rowsEffected int64, err Error)
}
