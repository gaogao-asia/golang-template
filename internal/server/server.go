package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/pkg/connection"
	slog "github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gin-gonic/gin"
)

func Run() {
	// create connection with database
	connection, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false
	api := engine.Group("/api")

	NewRouter(api, connection)

	// Listen HTTP Server
	srv := &http.Server{
		Addr:    config.AppConfig.Server.Port,
		Handler: engine,
	}

	go func() {
		// Start HTTP Server
		log.Printf("HTTP server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed to start: %v", err)
		}
	}()

	// Waiting signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	slog.InfoCtxf(ctx, "Graceful shutdown...")
	if err := srv.Shutdown(ctx); err != nil {
		slog.Fatalf("Graceful shutdown... err: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()

	slog.Infof("Graceful shutdown: good")
}
