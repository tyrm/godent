package terms

import (
	"context"

	"github.com/tyrm/godent/internal/logic"
)

// Module contains the terms module for the web server. Implements web.Module.
type Module struct {
	logic *logic.Logic
}

// New returns a new webapp module.
func New(_ context.Context, logicMod *logic.Logic) (*Module, error) {
	return &Module{
		logic: logicMod,
	}, nil
}

// Name return the module name.
func (*Module) Name() string {
	return "terms"
}
