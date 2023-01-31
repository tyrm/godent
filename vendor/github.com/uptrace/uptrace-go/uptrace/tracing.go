package uptrace

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"

	"github.com/uptrace/uptrace-go/internal"
)

func configureTracing(ctx context.Context, client *client, conf *config) {
	provider := conf.tracerProvider
	if provider == nil {
		var opts []sdktrace.TracerProviderOption

		opts = append(opts, sdktrace.WithIDGenerator(defaultIDGenerator()))
		if res := conf.newResource(); res != nil {
			opts = append(opts, sdktrace.WithResource(res))
		}
		if conf.traceSampler != nil {
			opts = append(opts, sdktrace.WithSampler(conf.traceSampler))
		}

		provider = sdktrace.NewTracerProvider(opts...)
		otel.SetTracerProvider(provider)
	}

	exp, err := otlptrace.New(ctx, otlpTraceClient(conf, client.dsn))
	if err != nil {
		internal.Logger.Printf("otlptrace.New failed: %s", err)
		return
	}

	queueSize := queueSize()
	bspOptions := []sdktrace.BatchSpanProcessorOption{
		sdktrace.WithMaxQueueSize(queueSize),
		sdktrace.WithMaxExportBatchSize(queueSize),
		sdktrace.WithBatchTimeout(10 * time.Second),
	}
	bspOptions = append(bspOptions, conf.bspOptions...)

	bsp := sdktrace.NewBatchSpanProcessor(exp, bspOptions...)
	provider.RegisterSpanProcessor(bsp)

	if conf.prettyPrint {
		exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			internal.Logger.Printf(err.Error())
		} else {
			provider.RegisterSpanProcessor(sdktrace.NewSimpleSpanProcessor(exporter))
		}
	}

	client.tp = provider
}

func otlpTraceClient(conf *config, dsn *DSN) otlptrace.Client {
	options := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(dsn.OTLPHost()),
		otlptracegrpc.WithHeaders(map[string]string{
			// Set the Uptrace DSN here or use UPTRACE_DSN env var.
			"uptrace-dsn": dsn.String(),
		}),
		otlptracegrpc.WithCompressor(gzip.Name),
	}

	if conf.tlsConf != nil {
		creds := credentials.NewTLS(conf.tlsConf)
		options = append(options, otlptracegrpc.WithTLSCredentials(creds))
	} else if dsn.Scheme == "https" {
		// Create credentials using system certificates.
		creds := credentials.NewClientTLSFromCert(nil, "")
		options = append(options, otlptracegrpc.WithTLSCredentials(creds))
	} else {
		options = append(options, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.NewClient(options...)
}

func queueSize() int {
	const min = 1000
	const max = 16000

	n := (runtime.GOMAXPROCS(0) / 2) * 1000
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}

//------------------------------------------------------------------------------

type idGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

var _ sdktrace.IDGenerator = (*idGenerator)(nil)

// NewIDs returns a new trace and span ID.
func (gen *idGenerator) NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	unixNano := time.Now().UnixNano()

	gen.Lock()
	defer gen.Unlock()

	tid := trace.TraceID{}
	binary.BigEndian.PutUint64(tid[:8], uint64(unixNano))
	_, _ = gen.randSource.Read(tid[8:])

	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])

	return tid, sid
}

// NewSpanID returns a ID for a new span in the trace with traceID.
func (gen *idGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID {
	gen.Lock()
	defer gen.Unlock()
	sid := trace.SpanID{}
	_, _ = gen.randSource.Read(sid[:])
	return sid
}

func defaultIDGenerator() *idGenerator {
	gen := &idGenerator{}
	var rngSeed int64
	_ = binary.Read(cryptorand.Reader, binary.LittleEndian, &rngSeed)
	gen.randSource = rand.New(rand.NewSource(rngSeed))
	return gen
}
