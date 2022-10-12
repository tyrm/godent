package generate

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/tyrm/godent/cmd/godent/action"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/logic"
)

// SigningKey generates a new signing key.
var SigningKey action.Action = func(ctx context.Context) error {
	logicMod := logic.New(
		nil,
		nil,
	)

	priv, err := logicMod.GenerateSigningKey()
	if err != nil {
		return err
	}

	privStr := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(priv.Seed())
	fmt.Printf("%s: \"ed25519 0 %s\"\n", config.Keys.SigningKey, privStr)

	return nil
}
