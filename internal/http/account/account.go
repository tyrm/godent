package account

import (
	"encoding/json"
	gdhttp "github.com/tyrm/godent/internal/http"
	"net/http"
)

type accountGetResponse struct {
	UserID string `json:"user_id"`
}

// accountGetHandler registers an account.
func (m *Module) accountGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "accountGetHandler")

	token, errCode, errString := m.logic.RequireAuth(r)
	if errCode != gdhttp.ErrCodeSuccess {
		gdhttp.ReturnError(w, errCode, errString)

		return
	}

	resp := accountGetResponse{
		UserID: token.Account.UserID,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}

// accountOptionsHandler registers nothing.
func (m *Module) accountOptionsHandler(_ http.ResponseWriter, _ *http.Request) {}
