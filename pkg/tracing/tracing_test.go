package tracing

import (
	"context"
	"testing"

	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/propagation"
)

func TestGetCarrierFromContext(t *testing.T) {
	propagator = propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})

	tests := []struct {
		name string
		ctx  context.Context
		want propagation.MapCarrier
	}{
		{
			name: "Input correct, get value",
			ctx:  context.Background(),
			want: propagation.MapCarrier{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := GetCarrierFromContext(tt.ctx)
			assert.Equal(t, tt.want, act)
		})
	}
}

func TestGetContextFromCarrier(t *testing.T) {
	propagator = propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})

	tests := []struct {
		name    string
		carrier propagation.MapCarrier
		want    context.Context
	}{
		{
			name:    "input nil, get nil",
			carrier: propagation.MapCarrier{},
			want:    context.Background(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := GetContextFromCarrier(tt.carrier)
			assert.Equal(t, tt.want, ctx)
		})
	}
}

func TestSpanStopStop(t *testing.T) {
	log.InitDev()
	// Create a new instance of the SpanStop struct with a mock span
	ctx := context.Background()

	s := SpanStop{
		FunctionName: "TestSpanStopStop",
	}

	// Call the Stop method
	s.End(ctx, nil)
}
