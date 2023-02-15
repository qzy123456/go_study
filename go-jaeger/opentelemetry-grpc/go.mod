module otelgrpc

go 1.18

require (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.39.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/exporters/jaeger v1.13.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0
	go.opentelemetry.io/otel/sdk v1.13.0
	go.opentelemetry.io/otel/trace v1.13.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require golang.org/x/net v0.6.0 // indirect
