[logging](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/logging)

## Enabling Extra Logging
We can enable logs and traces do diagnose the problem of your gRPC application. In the gRPC Go application, we can enable extra logs by setting the following environment variables:
```bash
# Verbosity means how many times any single info mesage should print every five minutes. The verbosity is set to 0 by default.
GRPC_GO_LOG_VERBOSITY_LEVEL=99 

# Sets log severity level to info. All the infomational messages will be printed.
GRPC_GO_LOG_SEVERITY_LEVEL=info 
```