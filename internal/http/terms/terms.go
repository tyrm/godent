package terms

import (
	"encoding/json"
	"net/http"
)

// TermsGetHandler returns the terms.
func (m *Module) TermsGetHandler(w http.ResponseWriter, _ *http.Request) {
	l := logger.WithField("func", "TermsGetHandler")

	terms := m.logic.GetTerms()

	err := json.NewEncoder(w).Encode(&terms)
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
