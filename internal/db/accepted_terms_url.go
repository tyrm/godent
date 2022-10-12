package db

import (
	"context"

	"github.com/tyrm/godent/internal/models"
)

type AcceptedTermsURL interface {
	CreateAcceptedTermsURL(ctx context.Context, acceptedTermsURLs ...*models.AcceptedTermsURL) Error
	ReadAcceptedTermsURLForAccount(ctx context.Context, accountID int64) ([]*models.AcceptedTermsURL, Error)
}
