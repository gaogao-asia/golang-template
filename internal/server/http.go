package server

import (
	"net/http"

	"github.com/gaogao-asia/golang-template/internal/server/http/handler"
	"github.com/gaogao-asia/golang-template/internal/server/http/middleware"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"github.com/gin-gonic/gin"
)

func NewRouter(api *gin.RouterGroup, connection connection.Conn) {
	// Middleware
	api.Use(middleware.CustomRecovery())
	api.Use(allowAllOrigins())

	// Routers
	v1 := handler.NewV1(api, connection)
	v1.Register()
}

func allowAllOrigins() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Request-ID")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
