package fc

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

func (c *Client) getHomeServer(ctx context.Context, serverName string) string {
	ctx, tracer := c.tracer.Start(ctx, "getHomeServer", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	cached := c.homeServerCache.Get(serverName)
	if cached != nil {
		return cached.Value()
	}

	// try to get http
	homeServer, cachePeriod, err := c.fetchServerWellKnown(ctx, serverName)
	if err == nil {
		c.homeServerCache.Set(serverName, homeServer, time.Duration(cachePeriod)*time.Second)

		return homeServer
	}

	// try to get dns
	homeServer, cachePeriod, err = c.fetchServerSRV(ctx, serverName)
	if err == nil {
		c.homeServerCache.Set(serverName, homeServer, time.Duration(cachePeriod)*time.Second)

		return homeServer
	}
	c.homeServerCache.Set(serverName, serverName, homeServerInvalidCachePeriod)

	return serverName
}
