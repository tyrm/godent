package fc

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/url"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

type serverWellKnown struct {
	MatrixServer string `json:"m.server"`
}

func (c *Client) fetchServerWellKnown(ctx context.Context, serverName string) (string, int, error) {
	ctx, tracer := c.tracer.Start(ctx, "fetchServerWellKnown", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	wellKnowURL := url.URL{
		Scheme: "https",
		Host:   serverName,
		Path:   "/.well-known/matrix/server",
	}

	resp, err := c.get(ctx, wellKnowURL.String())
	if err != nil {
		return "", 0, fmt.Errorf("get: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		return "", 0, ErrNotOKStatus
	}

	var data serverWellKnown
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", 0, fmt.Errorf("decode: %s", err.Error())
	}
	if data.MatrixServer == "" {
		return "", 0, ErrHomeServerNotFound
	}

	cachePeriod, err := cachePeriodFromHeaders(resp)
	if err != nil {
		cachePeriod = int(wellKnownDefaultCachePeriod.Seconds()) + rand.Intn(int(wellKnownDefaultCachePeriodJitter.Seconds()))
	} else {
		cachePeriod = int(math.Min(float64(cachePeriod), wellKnownDefaultMaxPeriod.Seconds()))
	}

	return data.MatrixServer, cachePeriod, nil
}

func (c *Client) fetchServerSRV(ctx context.Context, serverName string) (string, int, error) {
	_, tracer := c.tracer.Start(ctx, "fetchServerWellKnown", trace.WithSpanKind(trace.SpanKindInternal))
	defer tracer.End()

	_, srvs, err := net.LookupSRV("matrix", "tcp", serverName)
	if err != nil {
		return "", 0, fmt.Errorf("lookup: %s", err.Error())
	}
	if len(srvs) == 0 {
		return "", 0, ErrHomeServerNotFound
	}

	return fmt.Sprintf("%s:%d", strings.TrimSuffix(srvs[0].Target, "."), srvs[0].Port), int(dnsDefaultCachePeriod.Seconds()), nil
}
