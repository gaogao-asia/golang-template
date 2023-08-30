package middleware

import (
	"context"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(ctx context.Context) {
			if r := recover(); r != nil {
				functionName, filePath, line, _ := runtime.Caller(3)
				funcName := runtime.FuncForPC(functionName).Name()
				log.Printf("%s:%d:%s 	panic recover: %v", string(filePath), line, funcName, r)
				c.Abort()
			}
		}(c.Request.Context())

		c.Next()
	}
}
