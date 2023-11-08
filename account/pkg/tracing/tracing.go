package tracing

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lithammer/shortuuid"
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

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/pkg/log"
)

var Tracer trace.Tracer
var propagator propagation.TextMapPropagator
var conn *grpc.ClientConn

// InitTracing set trace base on tempo, loki via tempo otlp grpc port.
//
// - `gprcOTLPEndpoint` is tempo otlp grpc port without protocol, ex: localhost:4317
//
// - `serviceName“ is name of service, ex: user-service
func InitTracing() func() {
	if config.AppConfig.Monitor.OpenTelemetry.Enable {
		return initOpentelemetry()
	}

	return func() {}
}

func initOpentelemetry() func() {
	var err error
	var ctx = context.Background()
	conn, err = grpc.DialContext(ctx, config.AppConfig.Monitor.Tempo.Endpoint,
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
			semconv.ServiceNameKey.String(config.AppConfig.Server.Name),
		)),
	)

	otel.SetTracerProvider(tp)

	Tracer = otel.Tracer(config.AppConfig.Server.Name)
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

func GetStringCarrierFromCtx(ctx context.Context) string {
	carrier := GetCarrierFromContext(ctx)
	bytes, _ := json.Marshal(carrier)

	return string(bytes)
}

func GetContextFromCarrier(carrier propagation.MapCarrier) context.Context {
	if carrier == nil {
		return context.Background()
	}

	return propagator.Extract(context.Background(), carrier)
}

func GetContextFromStringCarrier(carrierStr string) context.Context {
	if carrierStr == "" {
		return context.Background()
	}
	var carrier propagation.MapCarrier
	_ = json.Unmarshal([]byte(carrierStr), &carrier)
	return GetContextFromCarrier(carrier)
}

func Start(ctx context.Context, params map[string]interface{}) (context.Context, SpanStop) {
	caller := log.GetFunctionNameAtRuntime(2)

	if params == nil {
		log.InfoCtxNoFuncf(ctx, "Start %s", caller.FunctionName)
	} else {
		prs := log.ToJsonString(params)
		log.InfoCtxNoFuncf(ctx, "Start %s, Function params: %+v", caller.FunctionName, prs)
	}

	if config.AppConfig.Monitor.OpenTelemetry.Enable {
		return start(ctx, caller.FunctionName)
	}

	return startWithoutOpentelemetry(ctx, caller.FunctionName)
}

func start(ctx context.Context, functionName string) (context.Context, SpanStop) {
	if conn.GetState() == connectivity.Ready {
		ctx, span := Tracer.Start(ctx, functionName)
		return ctx, SpanStop{span: span, FunctionName: functionName}
	}

	span := trace.SpanFromContext(ctx)

	return ctx, SpanStop{span: span, FunctionName: functionName}
}

func InitStartMiddleware(ctx context.Context, functionName string) (context.Context, SpanStop) {
	if conn.GetState() == connectivity.Ready {
		ctx, span := Tracer.Start(ctx, functionName)
		return ctx, SpanStop{span: span, FunctionName: functionName}
	}

	span := trace.SpanFromContext(ctx)

	return ctx, SpanStop{span: span, FunctionName: functionName}
}

func startWithoutOpentelemetry(ctx context.Context, functionName string) (context.Context, SpanStop) {
	span := newSpan(ctx, functionName)
	return ctx, span
}

type SpanStop struct {
	span         trace.Span
	StartTime    time.Time
	EndTime      time.Time
	FunctionName string
}

func (s SpanStop) End(ctx context.Context, params map[string]interface{}) {
	if s.span != nil {
		s.span.End()
	}

	// log all params with key-value format
	prs := log.ToJsonString(params)
	log.InfoCtxNoFuncf(ctx, "End: %s, Function result:%+v, ", s.FunctionName, prs)
}

func (s SpanStop) GetTraceID() string {
	if s.span != nil {
		return s.span.SpanContext().TraceID().String()
	}
	return shortuuid.New()
}

func newSpan(ctx context.Context, functionName string) SpanStop {
	return SpanStop{
		StartTime:    time.Now(),
		FunctionName: functionName,
		span:         nil,
	}
}
