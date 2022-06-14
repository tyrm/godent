package versions

import (
	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
	"net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *gdhttp.Server) error {
	s.HandleFunc(path.Versions, m.VersionGetHandler).Methods(http.MethodGet)

	return nil
}