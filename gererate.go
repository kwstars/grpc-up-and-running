package main

//go:generate protoc -I. -I$GOPATH/src --go_out=paths=source_relative:. --go-grpc_out=. --go-grpc_opt=paths=source_relative api/product_info/v1/product_info.proto
//go:generate protoc -I. -I$GOPATH/src --go_out=paths=source_relative:. --go-grpc_out=. --go-grpc_opt=paths=source_relative api/product_info/v2/product_info.proto
//go:generate protoc -I. -I$GOPATH/src --go_out=paths=source_relative:. --go-grpc_out=. --go-grpc_opt=paths=source_relative api/order_management/order_management.proto
