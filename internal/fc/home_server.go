package fc

import (
	"context"
	"errors"

	"github.com/tyrm/godent/internal/cache"
	"go.opentelemetry.io/otel/trace"
)

func (c *Client) getHomeServer(ctx context.Context, serverName string) string {
	ctx, tracer := c.tracer.Start(ctx, "getHomeServer", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	l := logger.WithField("func", "getHomeServer")

	homeServer, err := c.cache.GetHomeServer(ctx, serverName)
	if err != nil && !errors.Is(err, cache.ErrMiss) {
		l.Warnf("cache get: %s", err.Error())
	}
	if err == nil {
		return homeServer
	}

	// try to get http
	homeServer, cachePeriod, err := c.fetchServerWellKnown(ctx, serverName)
	if err == nil {
		cerr := c.cache.SetHomeServer(ctx, serverName, homeServer, cachePeriod)
		if cerr != nil {
			l.Warnf("cache set: %s", cerr.Error())
		}

		return homeServer
	}

	// try to get dns
	homeServer, _, err = c.fetchServerSRV(ctx, serverName)
	if err == nil {
		cerr := c.cache.SetHomeServer(ctx, serverName, homeServer, dnsDefaultCachePeriod)
		if cerr != nil {
			l.Warnf("cache set: %s", cerr.Error())
		}

		return homeServer
	}
	cerr := c.cache.SetHomeServer(ctx, serverName, homeServer, homeServerInvalidCachePeriod)
	if cerr != nil {
		l.Warnf("cache set: %s", cerr.Error())
	}

	return serverName
}
