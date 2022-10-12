package logic

import (
	"errors"
	"net/http"

	"github.com/tyrm/godent/internal/models"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db"
	gdhttp "github.com/tyrm/godent/internal/http"
	"go.opentelemetry.io/otel/trace"
)

func (logic *Logic) RequireAuth(r *http.Request) (*models.Account, gdhttp.ErrCode, string) {
	ctx, tracer := logic.tracer.Start(r.Context(), "RequireAuth", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	token, err := logic.tokenFromRequest(r)
	if err != nil {
		return nil, gdhttp.ErrCodeUnauthorized, gdhttp.ResponseUnauthorized
	}

	account, err := logic.db.ReadAccountByToken(ctx, token)
	if err != nil {
		if errors.Is(err, db.ErrNoEntries) {
			return nil, gdhttp.ErrCodeUnauthorized, gdhttp.ResponseUnauthorized
		}
		tracer.RecordError(err)

		return nil, gdhttp.ErrCodeUnknown, gdhttp.ResponseDatabaseError
	}

	if viper.GetBool(config.Keys.RequireTermsAgreed) {
		terms := logic.GetTerms(ctx)

		mv, err := terms.GetMasterVersion()
		if err == nil && account.ConsentVersion != mv {
			return nil, gdhttp.ErrCodeTermsNotSigned, gdhttp.ResponseTermsNotSigned
		}
	}

	return account, gdhttp.ErrCodeSuccess, ""
}
