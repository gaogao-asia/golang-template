package middleware

import (
	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
	"github.com/gin-gonic/gin"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var span tracing.SpanStop
		if config.AppConfig.Monitor.OpenTelemetry.Enable {
			ctx, span = tracing.InitStartMiddleware(c.Request.Context(), "TracingMiddleware")
			defer span.End(ctx, nil)
		}
		ctx = log.InitTraceIntoContext(ctx)

		r := c.Request.WithContext(ctx)
		c.Request = r

		log.InfoCtxNoFuncf(ctx, "Start TracingMiddleware")

		c.Next()
	}
}
