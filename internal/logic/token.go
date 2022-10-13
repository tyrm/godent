package logic

import (
	"context"

	"github.com/tyrm/godent/internal/models"
)

type Token interface {
	IssueToken(ctx context.Context, mxID string) (*models.Token, error)
}
