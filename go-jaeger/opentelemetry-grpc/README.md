# gRPC Tracing Example 地址

```$xslt
https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/google.golang.org/grpc/otelgrpc/example 
```
### 进入到 api目录，执行proto命令

```sh
# protobuf v1.3.2
protoc  --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.  hello-service.proto
```

###  server1

```sh
go run ./server
```

###  client1

```sh
go run ./client
```