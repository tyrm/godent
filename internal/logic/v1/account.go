package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/models"
	"go.opentelemetry.io/otel/trace"
)

func (logic *Logic) getOrCreateAccount(ctx context.Context, mxID string) (*models.Account, error) {
	ctx, tracer := logic.tracer.Start(ctx, "getOrCreateAccount", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	// try to get account
	account, err := logic.db.ReadAccountByUserID(ctx, mxID)
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		err = fmt.Errorf("read: %s", err.Error())
		tracer.RecordError(err)

		return nil, err
	}
	if err == nil {
		return account, nil
	}

	// create new account
	newAccount := &models.Account{
		UserID: mxID,
	}
	err = logic.db.CreateAccount(ctx, newAccount)
	if err != nil {
		err = fmt.Errorf("create: %s", err.Error())
		tracer.RecordError(err)

		return nil, err
	}

	return newAccount, nil
}
