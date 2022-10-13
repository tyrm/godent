package status

import (
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *gdhttp.Server) error {
	s.HandleFunc(path.V2, m.StatusGetHandler).Methods(http.MethodGet)
	s.HandleFunc(path.V2, m.StatusOptionsHandler).Methods(http.MethodOptions)

	return nil
}
