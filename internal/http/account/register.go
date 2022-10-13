package account

import (
	"encoding/json"
	"net/http"
	"strings"

	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/util"
)

type registerPostResponse struct {
	AccessToken string `json:"access_token"`
	Token       string `json:"token"`
}

// RegisterPostHandler registers an account.
func (m *Module) registerPostHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "RegisterPostHandler")

	// get request values
	accessToken := r.URL.Query().Get(gdhttp.QueryAccessToken)
	matrixServerName := r.URL.Query().Get(gdhttp.QueryMatrixServerName)

	if !util.IsValidMatrixServerName(matrixServerName) {
		gdhttp.ReturnError(w, gdhttp.ErrCodeInvalidParam, "matrix_server_name must be a valid Matrix server name (IP address or hostname)")

		return
	}

	userInfo, err := m.fc.OpenidUserinfo(r.Context(), matrixServerName, accessToken)
	if err != nil {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, err.Error())

		return
	}
	if userInfo.Sub == "" {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, "The Matrix homeserver did not include 'sub' in its response")

		return
	}

	userIDParts := strings.Split(userInfo.Sub, ":")
	if len(userIDParts) != 2 || !util.IsValidMatrixServerName(userIDParts[1]) {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, "The Matrix homeserver returned an invalid MXID")

		return
	}
	if userIDParts[1] != matrixServerName {
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, "The Matrix homeserver returned a MXID belonging to another homeserver")

		return
	}

	token, err := m.logic.IssueToken(r.Context(), userInfo.Sub)
	if err != nil {
		l.Errorf("issue token: %s", err.Error())
		gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, gdhttp.ResponseDatabaseError)

		return
	}

	resp := registerPostResponse{
		AccessToken: token.Token,
		Token:       token.Token,
	}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}

// RegisterOptionsHandler registers nothing.
func (m *Module) registerOptionsHandler(_ http.ResponseWriter, _ *http.Request) {}
