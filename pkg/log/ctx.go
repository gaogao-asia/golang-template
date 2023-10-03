package log

import (
	"context"

	"github.com/lithammer/shortuuid"
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

func GetTraceIDFromContext(ctx context.Context) string {
	if pre, ok := ctx.Value(CtxLogPrefix).(LogPrefix); ok {
		return pre.TraceID
	}

	return shortuuid.New()
}

func AddTraceIntoContext(ctx context.Context, traceID string) context.Context {
	pre := GetLogPrefix(ctx)
	if pre.TraceID != "" {
		return ctx
	}

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
