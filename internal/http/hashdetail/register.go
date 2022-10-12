package hashdetail

import (
	"encoding/json"
	"net/http"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/util"
)

// RegisterPostHandler registers an account.
func (m *Module) RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "RegisterPostHandler")

	// get request values
	// accessToken := r.URL.Query().Get(gdhttp.QueryAccessToken
	matrixServerName := r.URL.Query().Get(gdhttp.QueryMatrixServerName)

	if !util.IsValidMatrixServerName(matrixServerName) {
		gdhttp.ReturnError(w, gdhttp.ErrCodeInvalidParam, "matrix_server_name must be a valid Matrix server name (IP address or hostname)")

		return
	}

	err := json.NewEncoder(w).Encode(&struct{}{})
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
