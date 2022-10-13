package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/tyrm/godent/internal/logic"
	"github.com/tyrm/godent/internal/models"
	"github.com/tyrm/godent/internal/util"
	"go.opentelemetry.io/otel/trace"

	gdhttp "github.com/tyrm/godent/internal/http"
)

func (logic *Logic) DeleteToken(ctx context.Context, token *models.Token) error {
	return logic.db.DeleteToken(ctx, token)
}

func (logic *Logic) IssueToken(ctx context.Context, mxID string) (*models.Token, error) {
	ctx, tracer := logic.tracer.Start(ctx, "IssueToken", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	account, err := logic.getOrCreateAccount(ctx, mxID)
	if err != nil {
		err = fmt.Errorf("account: %s", err.Error())
		tracer.RecordError(err)

		return nil, err
	}

	newToken := &models.Token{
		Account:   account,
		AccountID: account.ID,
		Token:     util.RandString(64),
	}
	err = logic.db.CreateToken(ctx, newToken)
	if err != nil {
		err = fmt.Errorf("token: %s", err.Error())
		tracer.RecordError(err)

		return nil, err
	}

	return newToken, nil
}

func (*Logic) tokenFromRequest(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer "), nil
	}

	token = r.URL.Query().Get(gdhttp.QueryAccessToken)
	if token != "" {
		return token, nil
	}

	return "", logic.ErrNotFound
}
