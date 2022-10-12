package logic

import (
	"context"
	"go.opentelemetry.io/otel"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/http/fc"
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
	terms Terms
}

func (logic *Logic) Start(_ context.Context) error {
	l := logger.WithField("func", "Start")
	l.Info("starting logic")

	l.Debug("generating terms")
	logic.terms = genTerms()

	return nil
}
