package v1

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"github.com/tyrm/godent/internal/fc"
	"github.com/tyrm/godent/internal/models"
	"go.opentelemetry.io/otel"

	"github.com/tyrm/godent/internal/db"
	"go.opentelemetry.io/otel/trace"
)

// New returns a new logic module.
func New(db db.DB, fc *fc.Client) *Logic {
	// Since some functions in Logic can be used without initialization do not do anything here.
	// Simply return object and use Start() to initialize the module.
	return &Logic{
		db:     db,
		fc:     fc,
		tracer: otel.Tracer("internal/logic"),
	}
}

type Logic struct {
	db     db.DB
	fc     *fc.Client
	tracer trace.Tracer

	// generated values
	signingKey ed25519.PrivateKey
	terms      models.Terms
}

func (logic *Logic) Start(_ context.Context) error {
	l := logger.WithField("func", "Start")
	l.Info("starting logic")

	l.Debug("generating terms")
	logic.terms = genTerms()

	l.Debug("parsing private key")
	pk, err := getPrivateKey()
	if err != nil {
		return fmt.Errorf("private key: %s")
	}
	logic.signingKey = pk

	return nil
}
