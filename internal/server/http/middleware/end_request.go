package middleware

import (
	"bufio"
	"bytes"

	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gin-gonic/gin"
)

func EndRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Buffer the writer
		writer := bufio.NewWriter(c.Writer)
		c.Writer = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer, writer: writer}

		c.Next()

		statusCode := c.Writer.Status()
		body := c.Writer.(*bodyLogWriter).body.String()

		log.InfoCtxNoFuncf(c.Request.Context(), "Start EndRequestMiddleware: status=%d body=%s", statusCode, string(body))

		writer.Flush()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	writer *bufio.Writer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.writer.Write(b)
}
