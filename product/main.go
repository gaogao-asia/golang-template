package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/gaogao-asia/golang-catalog/pb"
)

type product struct {
	Name string
	pb.UnimplementedProductServiceServer
}

func (s *product) GetLatestProduct(ctx context.Context, in *pb.LastProductRequest) (*pb.LastProductResponse, error) {
	if in.Id == 0 {
		return nil, errors.New("id is required")
	}

	return &pb.LastProductResponse{Name: "Welcome"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &product{})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
