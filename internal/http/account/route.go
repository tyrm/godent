package account

import (
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *gdhttp.Server) error {
	s.HandleFunc(path.V2Account, m.accountGetHandler).Methods(http.MethodGet)
	s.HandleFunc(path.V2Account, m.accountOptionsHandler).Methods(http.MethodOptions)
	s.HandleFunc(path.V2AccountLogout, m.logoutOptionsHandler).Methods(http.MethodOptions)
	s.HandleFunc(path.V2AccountLogout, m.logoutPostHandler).Methods(http.MethodPost)
	s.HandleFunc(path.V2AccountRegister, m.registerOptionsHandler).Methods(http.MethodOptions)
	s.HandleFunc(path.V2AccountRegister, m.registerPostHandler).Methods(http.MethodPost)

	return nil
}
