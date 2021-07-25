[Google API Documentation](https://cloud.google.com/service-infrastructure/docs/service-management/reference/rpc/google.api#http)
- Each mapping needs to specify a URL path template and an HTTP method. 
- The path template can contain one or more fields in the gRPC request message. But those fields should be nonrepeated fields with primitive types.
- Any fields in the request message that are not in the path template automatically become HTTP query parameters if there is no HTTP request body.
- Fields that are mapped to URL query parameters should be either a primitive type or a repeated primitive type or a nonrepeated message type.
- For a repeated field type in query parameters, the parameter can be repeated in the URL as ...?param=A&param=B.
- For a message type in query parameters, each field of the message is mapped to a separate parameter, such as ...?foo.a=A&foo.b=B&foo.c=C.

## Dependent packages
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```


## Testing

* Add a new product to the ProductInfo service.

```
$ curl -X POST http://localhost:8081/v1/product -d '{"name": "Apple", "description": "iphone7", "price": 699}'

"38e13578-d91e-11e9-819f-6c96cfe0687d"
```

* Get the existing product using ProductID

```
$ curl http://localhost:8081/v1/product/38e13578-d91e-11e9-819f-6c96cfe0687d

{"id":"38e13578-d91e-11e9-819f-6c96cfe0687d","name":"Apple","description":"iphone7","price":
```