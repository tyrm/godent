package pubkey

import (
	"encoding/json"
	"net/http"
)

type isValidGetResponse struct {
	Valid bool `json:"valid"`
}

// isValidGetHandler registers a public key.
func (m *Module) isValidGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "isValidGetHandler")

	resp := isValidGetResponse{}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
