package account

import (
	"encoding/json"
	gdhttp "github.com/tyrm/godent/internal/http"
	"net/http"
)

func (m *Module) logoutPostHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "logoutPostHandler")

	token, errCode, errString := m.logic.RequireAuth(r)
	if errCode != gdhttp.ErrCodeSuccess {
		gdhttp.ReturnError(w, errCode, errString)

		return
	}

	err := m.logic.DeleteToken(r.Context(), token)
	if err != nil {
		l.Errorf("delete: %s", err.Error())
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, gdhttp.ResponseDatabaseError)

		return
	}

	if err := json.NewEncoder(w).Encode(&struct{}{}); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}

func (m *Module) logoutOptionsHandler(_ http.ResponseWriter, _ *http.Request) {}
