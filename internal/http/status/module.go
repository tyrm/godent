package status

import (
	"context"

	gdhttp "github.com/tyrm/godent/internal/http"
)

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct{}

// New returns a new webapp module.
func New(_ context.Context) (*Module, error) {
	return &Module{}, nil
}

// Name return the module name.
func (*Module) Name() string {
	return "status"
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(_ *gdhttp.Server) {}
