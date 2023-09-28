package middleware

import (
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
	"github.com/gin-gonic/gin"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check RequestID header
		ctx, span := tracing.Start(c.Request.Context(), nil)
		defer span.End(ctx, nil)
		traceID := span.GetTraceID()

		logPrefix := log.AddTraceIntoContext(ctx, traceID)
		r := c.Request.WithContext(logPrefix)
		c.Request = r

		c.Next()
	}
}
