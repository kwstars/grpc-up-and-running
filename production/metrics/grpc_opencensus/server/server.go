package main

import (
	"context"
	"errors"
	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	pb "github.com/kwstars/grpc-up-and-running/api/product_info/v2"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
	// Start z-Pages server.
	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe("127.0.0.1:8081", mux))
	}()

	// Register stats and tracing exporters to export
	// the collected data.
	view.RegisterExporter(&exporter.PrintExporter{})

	// Register the views to collect server request count.
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	// Create a gRPC Server with stats handler.
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	pb.RegisterProductInfoServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
