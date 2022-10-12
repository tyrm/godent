package logic

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/db"
	gdhttp "github.com/tyrm/godent/internal/http"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func (logic *Logic) RequireAuth(w http.ResponseWriter, r *http.Request) (gdhttp.ErrCode, string) {
	ctx, tracer := logic.tracer.Start(r.Context(), "RequireAuth", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	token, err := logic.tokenFromRequest(r)
	if err != nil {
		return gdhttp.ErrCodeUnauthorized, ResponseUnauthorized
	}

	account, err := logic.db.ReadAccountByToken(ctx, token)
	if err != nil {
		if errors.Is(err, db.ErrNoEntries) {
			return gdhttp.ErrCodeUnauthorized, ResponseUnauthorized
		}
		tracer.RecordError(err)

		return gdhttp.ErrCodeUnknown, ResponseDatabaseError
	}

	_ = account

	if viper.GetBool(config.Keys.RequireTermsAgreed) {

	}

	return gdhttp.ErrCodeUnknown, ""
}
