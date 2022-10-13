package versions

import (
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *gdhttp.Server) error {
	s.HandleFunc(path.Versions, m.VersionGetHandler).Methods(http.MethodGet)

	return nil
}
