package logic

import (
	"context"
	"github.com/tyrm/godent/internal/models"
)

type Terms interface {
	AddAgreedTerms(ctx context.Context, account *models.Account, urls []string) error
	GetTerms(ctx context.Context) models.Terms
}
