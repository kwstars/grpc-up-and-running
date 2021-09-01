package main

//go:generate mockgen -destination production/grpc_continuous_integration/mock/prodinfo_mock.go -package mock github.com/kwstars/grpc-up-and-running/api/product_info/v2 ProductInfoClient
