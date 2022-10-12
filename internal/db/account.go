package db

import (
	"context"
	"github.com/tyrm/godent/internal/models"
)

type Account interface {
	ReadAccountByToken(ctx context.Context, token string) (*models.Account, error)
}
