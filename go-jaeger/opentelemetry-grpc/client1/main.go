package main

import (
	"otelgrpc/api"
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	service = "mxshop"
)

//traceProvider返回配置为使用的OpenTelemetry tracerProvider
// Jaeger导出器，它将向提供的url发送跨度。返回的 TracerProvider还将使用配置了所有信息的资源
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

	//conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),   // 拦截普通的一次请求一次响应的rpc服务
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()), // 拦截流式的rpc服务
	)
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
    //grpc执行调用
	c := api.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &api.HelloRequest{Greeting: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Reply)
}


