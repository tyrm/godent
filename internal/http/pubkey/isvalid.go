package pubkey

import (
	"encoding/json"
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
)

type isValidGetResponse struct {
	Valid bool `json:"valid"`
}

// isValidGetHandler registers a public key.
func (m *Module) isValidGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "isValidGetHandler")

	publicKey := r.URL.Query().Get(gdhttp.QueryPublicKey)
	valid, err := m.logic.IsPubKeyValid(r.Context(), publicKey)
	if err != nil {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, err.Error())

		return
	}

	resp := isValidGetResponse{
		Valid: valid,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
