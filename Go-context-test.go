package main
import (
	"context"
	"fmt"
	"time"
)

func Println(ctx context.Context, a, b int) {
	for {
		fmt.Println(a + b)
		a, b = a+1, b+1
		select {
		case <-ctx.Done():
			fmt.Println("程序结束")
			return
		default:
		}
	}
}

func main() {
	{
		// 超时取消
		a := 1
		b := 2
		timeout := 2 * time.Second
		ctxBg := context.Background()
		ctx, _ := context.WithTimeout(ctxBg, timeout)
		Println(ctx, a, b)

		time.Sleep(2 * time.Second) // 等待时候还会继续输出
	}
	{
		// 手动取消
		a := 1
		b := 2
		ctx, cancelCtx := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancelCtx() // 在调用处主动取消
		}()
		Println(ctx, a, b)

		time.Sleep(2 * time.Second)
	}
}