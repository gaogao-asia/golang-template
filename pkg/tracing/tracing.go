package tracing

import (
	"context"

	"github.com/gaogao-asia/golang-template/pkg/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdttrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var Tracer trace.Tracer
var propagator propagation.TextMapPropagator
var conn *grpc.ClientConn

// InitTracing set trace base on tempo, loki via tempo otlp grpc port.
//
// - `gprcOTLPEndpoint` is tempo otlp grpc port without protocol, ex: localhost:4317
//
// - `serviceName“ is name of service, ex: user-service
func InitTracing(ctx context.Context, gprcOTLPEndpoint, serviceName string) func() {
	var err error
	conn, err = grpc.DialContext(ctx, gprcOTLPEndpoint,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Errorf("InitTracing: failed to create gRPC connection to collector: %v", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Errorf("InitTracing: failed to initialize exporter: %v", err)
	}

	tp := sdttrace.NewTracerProvider(
		sdttrace.WithBatcher(exporter),
		sdttrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)

	Tracer = otel.Tracer(serviceName)
	propagator = propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})

	return func() {
		if err := tp.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}

func GetCarrierFromContext(ctx context.Context) propagation.MapCarrier {
	if ctx == nil {
		ctx = context.Background()
	}

	// Serialize the context into carrier
	carrier := propagation.MapCarrier{}
	propagator.Inject(ctx, carrier)

	return carrier
}

func GetContextFromCarrier(carrier propagation.MapCarrier) context.Context {
	if carrier == nil {
		return context.Background()
	}

	return propagator.Extract(context.Background(), carrier)
}

func Start(ctx context.Context, params map[string]interface{}) (context.Context, SpanStop) {
	caller := log.GetFunctionNameAtRuntime(2)

	if params == nil {
		log.InfoCtxNoFuncf(ctx, "Start %s", caller.FunctionName)
	} else {
		prs := log.ToJsonString(params)
		log.InfoCtxNoFuncf(ctx, "Start %s, Function params: %+v", caller.FunctionName, prs)
	}

	return start(ctx, caller)
}

func start(ctx context.Context, caller log.Caller) (context.Context, SpanStop) {
	if conn.GetState() == connectivity.Ready {
		ctx, span := Tracer.Start(ctx, caller.FunctionName)
		return ctx, SpanStop{span: span, caller: caller}
	}

	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()
	ctx = log.AddTraceIntoContext(ctx, traceID)

	return ctx, SpanStop{span: span, caller: caller}
}

type SpanStop struct {
	span   trace.Span
	caller log.Caller
}

func (s SpanStop) End(ctx context.Context, params map[string]interface{}) {
	if s.span != nil {
		s.span.End()
	}

	// log all params with key-value format
	prs := log.ToJsonString(params)
	log.InfoCtxNoFuncf(ctx, "End: %s, Function result:%+v, ", s.caller.FunctionName, prs)
}

func (s SpanStop) GetTraceID() string {
	return s.span.SpanContext().TraceID().String()
}
