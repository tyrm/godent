package logic

import (
	"context"

	"github.com/tyrm/godent/internal/db"
	"github.com/tyrm/godent/internal/http/fc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// New returns a new logic module.
func New(_ context.Context, db db.DB, fc *fc.Client) (*Logic, error) {
	return &Logic{
		db:     db,
		fc:     fc,
		tracer: otel.Tracer("internal/logic"),
	}, nil
}

type Logic struct {
	db     db.DB
	fc     *fc.Client
	tracer trace.Tracer
}
