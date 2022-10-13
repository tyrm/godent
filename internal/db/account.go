package db

import (
	"context"

	"github.com/tyrm/godent/internal/models"
)

type Account interface {
	CreateAccount(ctx context.Context, account *models.Account) Error
	// ReadAccountByToken(ctx context.Context, token string) (*models.Account, Error)
	ReadAccountByUserID(ctx context.Context, userID string) (*models.Account, Error)
	UpdateAccount(ctx context.Context, account *models.Account) Error
}
