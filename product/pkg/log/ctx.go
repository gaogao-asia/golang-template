package log

import (
	"context"

	"github.com/lithammer/shortuuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gaogao-asia/golang-catalog/config"
)

type CtxKey string

const (
	CtxUserIDKey  CtxKey = "userID"
	CtxTraceIDKey CtxKey = "traceID"
	CtxLogPrefix  CtxKey = "logPrefix"
	CtxTokenKey   CtxKey = "token"
)

type LogPrefix struct {
	UserID    string
	TraceID   string
	RequestID string
}

func GetUserIDFromContext(ctx context.Context) string {
	pre := GetLogPrefix(ctx)
	return pre.UserID
}

func AddUserIntoContext(ctx context.Context, userID string) context.Context {
	pre := GetLogPrefix(ctx)
	pre.UserID = userID
	return context.WithValue(ctx, CtxLogPrefix, pre)
}

func InitTraceIntoContext(ctx context.Context) context.Context {
	var traceID string

	if config.AppConfig.Monitor.OpenTelemetry.Enable {
		span := trace.SpanFromContext(ctx)
		traceID = span.SpanContext().TraceID().String()
	} else {
		traceID = shortuuid.New()
	}

	pre := GetLogPrefix(ctx)

	pre.TraceID = traceID

	return context.WithValue(ctx, CtxLogPrefix, pre)
}

func AddRequestIDIntoContext(ctx context.Context, requestID string) context.Context {
	pre := GetLogPrefix(ctx)
	pre.RequestID = requestID
	return context.WithValue(ctx, CtxLogPrefix, pre)
}

func GetLogPrefix(ctx context.Context) LogPrefix {
	if pre, ok := ctx.Value(CtxLogPrefix).(LogPrefix); ok {
		return pre
	}

	return LogPrefix{}
}

func AddTokenIntoContext(ctx context.Context, atoken string) context.Context {
	return context.WithValue(ctx, CtxTokenKey, atoken)
}

func GetTokenFromContext(ctx context.Context) string {
	if token, ok := ctx.Value(CtxTokenKey).(string); ok {
		return token
	}

	return ""
}
