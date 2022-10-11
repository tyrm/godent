package status

import (
	"encoding/json"
	"net/http"
)

// StatusGetHandler returns an empty object.
func (m *Module) StatusGetHandler(w http.ResponseWriter, _ *http.Request) {
	l := logger.WithField("func", "StatusGetHandler")

	err := json.NewEncoder(w).Encode(&struct{}{})
	if err != nil {
		l.Errorf("encoding version response: %s", err.Error())
	}
}
