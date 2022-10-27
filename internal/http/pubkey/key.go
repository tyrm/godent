package pubkey

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	gdhttp "github.com/tyrm/godent/internal/http"
	"github.com/tyrm/godent/internal/path"
)

type keyGetResponse struct {
	PublicKey string `json:"public_key"`
}

// keyGetHandler registers a public key.
func (m *Module) keyGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "keyGetHandler")

	keyID := mux.Vars(r)[path.VarKeyID]
	switch keyID {
	case "ed25519:0":
		pubStr, err := m.logic.GetPublicKey()
		if err != nil {
			gdhttp.ReturnError(w, gdhttp.ErrCodeUnknown, err.Error())

			return
		}

		resp := keyGetResponse{
			PublicKey: pubStr,
		}
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			l.Errorf("encoding response: %s", err.Error())
		}
	default:
		gdhttp.ReturnError(w, gdhttp.ErrCodeNotFound, "key not found")
	}
}
