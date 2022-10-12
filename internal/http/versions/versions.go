package versions

import (
	"encoding/json"
	"net/http"
)

var supportedVersions = []string{
	"r0.2.0",
	"r0.2.1",
	"v1.1",
}

type VersionGetResp struct {
	Versions []string `json:"versions"`
}

// VersionGetHandler gets the versions of the specification supported by the server.
func (m *Module) VersionGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "VersionGetHandler")

	resp := VersionGetResp{
		Versions: supportedVersions,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		l.Errorf("encoding response: %s", err.Error())
	}
}
