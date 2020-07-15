package main
import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("主协程结束。")

	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕。")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
		ch <- "子协程干活儿了。" //把这行注释掉再运行一下，看看什么结果(报错 fatal error: all goroutines are asleep - deadlock! 死锁)
		       					//当 channel是nil的时候，无论是传入数据还是取出数据，都会永久的block。
	}()

	str,ok := <-ch    //没有数据前，阻塞
	if ok {
		fmt.Println("这是判断channel",ok)
	}
	fmt.Println("str = ", str)
}