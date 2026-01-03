package main

import (
	"context"
	"log"
	"net"

	pb "grpc-mock/proto"

	"google.golang.org/grpc"
)

// Server implements the ProductService gRPC server
type Server struct {
	pb.UnimplementedProductServiceServer
	products []*pb.Product
}

// NewServer creates a new gRPC server instance with mock data
func NewServer() *Server {
	return &Server{
		products: []*pb.Product{
			{Id: 1, Name: "Laptop", Description: "High-performance laptop", Price: 1299.99, Stock: 50},
			{Id: 2, Name: "Mouse", Description: "Wireless mouse", Price: 29.99, Stock: 200},
			{Id: 3, Name: "Keyboard", Description: "Mechanical keyboard", Price: 89.99, Stock: 150},
			{Id: 4, Name: "Monitor", Description: "4K Monitor 27 inch", Price: 399.99, Stock: 75},
			{Id: 5, Name: "Headphones", Description: "Noise-cancelling headphones", Price: 249.99, Stock: 100},
		},
	}
}

// GetProducts returns all products
func (s *Server) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	log.Println("📦 GetProducts called")
	return &pb.GetProductsResponse{Products: s.products}, nil
}

// GetProductByID returns a specific product by ID
func (s *Server) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.Product, error) {
	log.Printf("📦 GetProductByID called: ID=%d", req.Id)

	for _, p := range s.products {
		if p.Id == req.Id {
			return p, nil
		}
	}

	return nil, nil
}

// CreateProduct creates a new product
func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	log.Printf("📦 CreateProduct called: Name=%s", req.Name)

	newProduct := &pb.Product{
		Id:          int32(len(s.products) + 1),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	s.products = append(s.products, newProduct)
	return newProduct, nil
}

func main() {
	listener, err := net.Listen("tcp", ":17003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, NewServer())

	log.Println("🚀 gRPC Mock Server starting on :17003")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
