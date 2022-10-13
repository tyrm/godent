package account

import (
	"context"

	"github.com/tyrm/godent/internal/fc"

	"github.com/tyrm/godent/internal/logic"
)

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	fc    *fc.Client
	logic logic.Logic
}

// New returns a new webapp module.
func New(_ context.Context, fc *fc.Client, logicMod logic.Logic) (*Module, error) {
	return &Module{
		fc:    fc,
		logic: logicMod,
	}, nil
}

// Name return the module name.
func (*Module) Name() string {
	return "account"
}
