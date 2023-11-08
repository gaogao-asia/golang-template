package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/gaogao-asia/golang-catalog/pb"

	"github.com/gaogao-asia/golang-template/config"
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
)

type ProductResponse struct {
	Name string
}

type ProductClient interface {
	GetLatestProduct(ctx context.Context, productID int) (*ProductResponse, error)
}

type productClient struct{}

func NewProductClient() ProductClient {
	return &productClient{}
}

func (p *productClient) GetLatestProduct(ctx context.Context, productID int) (res *ProductResponse, err error) {
	ctx, span := tracing.Start(ctx, log.Print{"productID": productID})
	defer span.End(ctx, log.Print{"product": &res})
	log.Infof("Context: %v", ctx)
	traceKey := tracing.GetStringCarrierFromCtx(ctx)

	conn, err := grpc.Dial(config.AppConfig.ProductClient.Endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductServiceClient(conn)

	result, err := c.GetLatestProduct(ctx, &pb.LastProductRequest{
		TraceKey: traceKey,
		Id:       int64(productID),
	})

	if err != nil {
		log.ErrorCtxf(ctx, "Call client get product error: %v", err)
		return nil, err
	}

	// connect to product service
	return &ProductResponse{result.GetName()}, nil
}
