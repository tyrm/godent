package logic

import (
	"net/http"
	"strings"
)

func (*Logic) tokenFromRequest(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer "), nil
	}

	token = r.URL.Query().Get(QueryAccessToken)
	if token != "" {
		return token, nil
	}

	return "", ErrTokenNotFound
}
