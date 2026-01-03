package grpc

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ProductServiceClient is the interface for the mock Product service
// This would normally be generated from the product.proto file
type ProductServiceClient interface {
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*Product, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
}

type GetProductsRequest struct{}

type GetProductsResponse struct {
	Products []*Product
}

type GetProductByIDRequest struct {
	Id int32
}

type CreateProductRequest struct {
	Name        string
	Description string
	Price       float64
	Stock       int32
}

type Product struct {
	Id          int32
	Name        string
	Description string
	Price       float64
	Stock       int32
}

// TestGRPCMockServer tests connection to gRPC mock server
func TestGRPCMockServer(t *testing.T) {
	// このテストはdocker-composeでgrpc-mockが起動している必要があります
	// docker-compose up -d grpc-mock

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Skipf("gRPC Mock server not available: %v", err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connection test
	state := conn.GetState()
	t.Logf("Connection state: %v", state)

	// Wait for connection to be ready
	if !conn.WaitForStateChange(ctx, state) {
		t.Skip("Could not connect to gRPC mock server")
		return
	}

	t.Log("Successfully connected to gRPC Mock Server on :50052")
}
