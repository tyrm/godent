package pubkey

import (
	"encoding/json"
	"net/http"
)

type ephemeralIsValidGetResponse struct {
	Valid bool `json:"valid"`
}

// ephemeralIsValidGetHandler registers a public key.
func (m *Module) ephemeralIsValidGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ephemeralIsValidGetHandler")

	resp := ephemeralIsValidGetResponse{}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
