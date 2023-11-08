package main

import (
	"context"
	"errors"
	"flag"
	"net"

	"google.golang.org/grpc"

	"github.com/gaogao-asia/golang-catalog/config"
	"github.com/gaogao-asia/golang-catalog/pb"
	"github.com/gaogao-asia/golang-catalog/pkg/log"
	"github.com/gaogao-asia/golang-catalog/pkg/tracing"
)

type product struct {
	Name string
	pb.UnimplementedProductServiceServer
}

func (s *product) GetLatestProduct(ctx context.Context, in *pb.LastProductRequest) (res *pb.LastProductResponse, err error) {
	if ctx != nil && ctx != context.Background() {
		log.Infof("Context: %v", ctx)
	}

	nctx := tracing.GetContextFromStringCarrier(in.TraceKey)
	log.Infof("Context: %v", ctx)

	nctx, span := tracing.Start(nctx, log.Print{"request param": in})
	defer span.End(nctx, log.Print{"response": &res})

	if in.Id == 0 {
		return nil, errors.New("id is required")
	}

	return &pb.LastProductResponse{Name: "Last lesson for golang"}, nil
}

func init() {
	configPath := flag.String("config", "./config", "config folder path")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	config.AppConfig = cfg

	log.InitDev()
	tracing.InitTracing()

	log.Infof("config: %+v", config.AppConfig)
}

func main() {
	lis, err := net.Listen("tcp", config.AppConfig.Server.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &product{})

	log.Infof("Server started on :%s", config.AppConfig.Server.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
