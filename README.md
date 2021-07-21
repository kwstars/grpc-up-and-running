## Quick start

**1. Install Go**

[Download and install](https://golang.org/doc/install)

**2. Install Protocol buffer**

[Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)

**3. Go plugins for the protocol compiler**

```bash
# Install the protocol compiler plugins for Go using the following commands:
go install google.golang.org/protobuf/advanced/protoc-gen-go@v1.26
go install google.golang.org/grpc/advanced/protoc-gen-go-grpc@v1.1

# Update your PATH so that the protoc compiler can find the plugins
export PATH="$PATH:$(go env GOPATH)/bin"
```


--- 
[Quick start for gRPC](https://grpc.io/docs/languages/go/quickstart/)

