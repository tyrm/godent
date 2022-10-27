package pubkey

import (
	"encoding/json"
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
)

type ephemeralIsValidGetResponse struct {
	Valid bool `json:"valid"`
}

// ephemeralIsValidGetHandler registers a public key.
func (m *Module) ephemeralIsValidGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ephemeralIsValidGetHandler")

	publicKey := r.URL.Query().Get(gdhttp.QueryPublicKey)
	valid, err := m.logic.IsEphemeralPubKeyValid(r.Context(), publicKey)
	if err != nil {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, err.Error())

		return
	}

	resp := ephemeralIsValidGetResponse{
		Valid: valid,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
