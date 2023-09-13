package middleware

import (
	"github.com/gaogao-asia/golang-template/internal/server/http/response"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gin-gonic/gin"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check RequestID header
		requestID := c.Request.Header.Get("X-Request-ID")
		// if requestID is empty, return badrequest
		if requestID == "" {
			body := response.ResponseBody{
				Error: &response.ErrorResponseBody{
					Code:    errs.ErrXRequestIDMissed.Code,
					Message: errs.ErrXRequestIDMissed.MsgCode,
				},
			}

			c.JSON(400, body)
			c.Abort()
			return
		}
		log.InfoCtxf(c.Request.Context(), "Start request: Add requestID: %s", requestID)

		logPrefix := log.AddRequestIDIntoContext(c.Request.Context(), requestID)
		r := c.Request.WithContext(logPrefix)
		c.Request = r

		c.Next()
	}
}
