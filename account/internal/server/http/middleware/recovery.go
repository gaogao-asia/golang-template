package middleware

import (
	"context"

	"github.com/gaogao-asia/golang-template/pkg/log"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(ctx context.Context) {
			if r := recover(); r != nil {
				caller := log.GetFunctionNameAtRuntime(6)
				log.ErrorCtxf(ctx, "panic recover: %s:%d:%s 	%v", caller.FilePath, caller.Line, caller.FunctionName, r)
				c.Abort()
			}
		}(c.Request.Context())

		c.Next()
	}
}
