package main

import (
	"otelgrpc/api"
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	service = "mxshop-srv"
)

type Server struct {
	api.UnimplementedHelloServiceServer // 这是 它内部实现的 的一个 结构体   我这边调用 以实现鸭子类型
}

func (s *Server) SayHello(ctx context.Context, request *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{
		Reply: "hello " + request.Greeting,
	}, nil
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// 创建 Jaeger 导出器
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
		)),
	)
	return tp, nil
}

func main() {
	// 用来创建 Jaeger 导出器
	tp, err := tracerProvider("http://192.168.16.51:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// 注册我们的 TracerProvider 为全局的 所以任何导入
	// 将来的 instrumentation 将默认使用它
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	// 创建 子 context  用来传递给 子协程  用于 通信 底层实现的是 channel   cancel 用于关闭用
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 当应用程序退出时，干净地关闭和刷新遥测
	defer func(ctx context.Context) {
		// 在关闭应用程序时，不要使其挂起
		ctx, cancel := context.WithTimeout(ctx, time.Second*5) // 从创建后 超过5秒后 关闭
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil { // Shutdown按注册span处理器的顺序关闭它们
			log.Fatal(err)
		}
	}(ctx)

	// 调用远程grpc
	g := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	api.RegisterHelloServiceServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
}


