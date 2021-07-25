package main

import (
	"context"
	"errors"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	pb "github.com/kwstars/grpc-up-and-running/api/product_info/v2"
	"github.com/kwstars/grpc-up-and-running/production/tracing/grpc_opentracing/tracer"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrapper.StringValue{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID" + in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create a gRPC Server with gRPC interceptor.
	grpcServer := NewServer()

	pb.RegisterProductInfoServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServer() *grpc.Server {
	// initialize jaegertracer
	jaegertracer, closer, err := tracer.NewTracer("product_mgt")
	if err != nil {
		log.Fatalln(err)
		return grpc.NewServer()
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(jaegertracer)

	// initialize grpc server with chained interceptors
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			// add opentracing unary interceptor to chain
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(jaegertracer)),
		),
	)
	return server
}
