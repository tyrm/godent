package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tyrm/godent/internal/models"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db"
	gdhttp "github.com/tyrm/godent/internal/http"
	"go.opentelemetry.io/otel/trace"
)

func (logic *Logic) RequireAuth(r *http.Request) (*models.Token, gdhttp.ErrCode, string) {
	ctx, tracer := logic.tracer.Start(r.Context(), "RequireAuth", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	l := logger.WithField("func", "RequireAuth")

	reqToken, err := logic.tokenFromRequest(r)
	if err != nil {
		l.Tracef("token from request: %s", err.Error())

		return nil, gdhttp.ErrCodeUnauthorized, gdhttp.ResponseUnauthorized
	}

	token, err := logic.db.ReadTokenByToken(ctx, reqToken)
	if err != nil {
		if errors.Is(err, db.ErrNoEntries) {
			l.Tracef("token not found")

			return nil, gdhttp.ErrCodeUnauthorized, gdhttp.ResponseUnauthorized
		}
		err = fmt.Errorf("read token: %s", err.Error())
		l.Tracef(err.Error())
		tracer.RecordError(err)

		return nil, gdhttp.ErrCodeUnknown, gdhttp.ResponseDatabaseError
	}

	if token.Account == nil {
		l.Tracef("account not found")

		return nil, gdhttp.ErrCodeUnauthorized, gdhttp.ResponseUnauthorized
	}

	if viper.GetBool(config.Keys.RequireTermsAgreed) {
		terms := logic.GetTerms(ctx)
		if terms.MasterVersion != "" && token.Account.ConsentVersion != terms.MasterVersion {
			l.Tracef("account hasn't signed terms")

			return nil, gdhttp.ErrCodeTermsNotSigned, gdhttp.ResponseTermsNotSigned
		}
	}

	return token, gdhttp.ErrCodeSuccess, ""
}
