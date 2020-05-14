package main
import (
	"fmt"
	"time"
)

func main()  {
	//创建一个无缓存的channel
	ch := make(chan int,0)

	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n",len(ch),cap(ch))

	//新建协程
	go func() {
		for i:=0;i<3;i++{
			fmt.Println("子协程：i=",i)
			ch <- i
		}
	}()

	//延时
	time.Sleep(2*time.Second)

	for i:=0;i<3;i++{
		num := <-ch    //读取管道中内容，没有内容前，阻塞
		fmt.Println("num =",num)
	}


	/////////////
	inCh := make(chan int)
	outCh := make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
				case out <- val:
					println("--------")
					out = nil
					in = inCh
				case val = <-in:
					println("++++++++++")
					out = outCh
					in = nil
			}
		}
	}()

	go func() {
		for r := range outCh {
			fmt.Println("Result: ", r)
		}
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}