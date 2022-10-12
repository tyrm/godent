package fc

import (
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func New(httpClient *http.Client) *Client {
	return &Client{
		http:   httpClient,
		tracer: otel.Tracer("internal/http/fc"),
	}
}

type Client struct {
	http   *http.Client
	tracer trace.Tracer
}
