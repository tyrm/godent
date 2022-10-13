package db

import (
	"context"

	"github.com/tyrm/godent/internal/models"
)

type Token interface {
	CreateToken(ctx context.Context, token *models.Token) Error
	ReadTokenByToken(ctx context.Context, token string) (*models.Token, Error)
	DeleteToken(ctx context.Context, token *models.Token) Error
}
