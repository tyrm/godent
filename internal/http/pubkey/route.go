package pubkey

import (
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *gdhttp.Server) error {
	s.HandleFunc(path.V2PubkeyEphemeralIsvalid, m.ephemeralIsValidGetHandler).Methods(http.MethodGet)
	s.HandleFunc(path.V2PubkeyIsvalid, m.isValidGetHandler).Methods(http.MethodGet)
	s.HandleFunc(path.V2PubkeyKey, m.keyGetHandler).Methods(http.MethodGet)

	return nil
}
