package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/internal/server"
	glog "github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
)

func main() {
	// Load configuration
	configPath := flag.String("config", "./config", "config folder path")
	flag.Parse()

	log.Printf("ENV: %s", os.Getenv("APPENV"))
	log.Printf("Config path: %s", *configPath)

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	config.AppConfig = cfg
	glog.InitDev()
	tracing.InitTracing(context.Background(), "tempo:4317", "golang-template")
	// Run server
	server.Run()
}
