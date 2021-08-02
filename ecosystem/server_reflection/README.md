## gRPC command line tool
https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md


## Testing

* List Services

```
$ ./grpc_cli ls localhost:50051

Output:
ecommerce.ProductInfo
grpc.reflection.v1alpha.ServerReflection
```

* List service details

```
$ ./grpc_cli ls localhost:50051 ecommerce.ProductInfo -l

Output:
package: ecommerce;
service ProductInfo {
rpc addProduct(ecommerce.Product) returns (google.protobuf.StringValue) {}
rpc getProduct(google.protobuf.StringValue) returns (ecommerce.Product) {}
}
```

* List method details

```
$ ./grpc_cli ls localhost:50051 ecommerce.ProductInfo.addProduct -l

Output:
rpc addProduct(ecommerce.Product) returns (google.protobuf.StringValue) {}
```server-reflection
```
$ ./grpc_cli type localhost:50051 ecommerce.Product

Output:
message Product {
string id = 1[json_name = "id"];
string name = 2[json_name = "name"];
string description = 3[json_name = "description"];
float price = 4[json_name = "price"];
}
```

* Call remote methods
```
$ ./grpc_cli call localhost:50051 addPrserver-reflection
value: "d962db94-d907-11e9-b49b-6c96cfe0687d"
Rpc succeeded with OK status
```

## Additional Information

### Generate Server and Client side code
```
protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/proto
```

### Update after changing the service definition
```
go get -u github.com/grpc-up-and-running/samples/ch08/server-reflection/go/proto
```
