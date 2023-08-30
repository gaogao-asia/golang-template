package server

import (
	"log"
	"net/http"

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/internal/server/http/handler"
	"github.com/gaogao-asia/golang-template/internal/server/http/middleware"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"github.com/gin-gonic/gin"
)

func newHTTPServer(connection connection.Conn) {
	// HTTP Server
	engine := gin.New()
	engine.RedirectTrailingSlash = false

	engine.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := engine.Group("/api")

	// Middleware
	api.Use(middleware.CustomRecovery())
	api.Use(allowAllOrigins())

	// Routers
	v1 := handler.NewV1(api, connection)
	v1.Register()

	// Listen HTTP Server
	srv := &http.Server{
		Addr:    config.AppConfig.Server.Port,
		Handler: engine,
	}

	log.Printf("HTTP server listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
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
