package client

import (
	"context"
)

type ProductResponse struct {
	Name string
}

type ProductClient interface {
	GetLatestProduct(ctx context.Context, productID int, traceKey string) (*ProductResponse, error)
}
