package v1

import (
	"github.com/tyrm/godent/internal/logic"
	"net/http"
	"strings"

	gdhttp "github.com/tyrm/godent/internal/http"
)

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
