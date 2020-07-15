package main
import (
	"fmt"
	"os"
	"sync"
	"time"
)
type Request struct {
	data []int
	ret chan int
}
func NewRequest(data ...int) *Request {
	return &Request{ data, make(chan int, 1) }
}
func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
	req.ret <- x
}
func main()  {
	var wg sync.WaitGroup
	quit := make(chan bool)

	for i:=0;i<2 ;i++  {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			task := func() {
				println(id,time.Now().Nanosecond())
				time.Sleep(time.Second)
			}
			for{
				select {
				case <-quit: //close channel 不会阻塞，因此可以用作退出通知
					return
				default:  //正常执行任务
					task()
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 3) //让测试goruntine运行一会,其实也就是会打印多少遍 print
	close(quit) //发出退出通知
	wg.Wait()
	///////////////////////////////
	//channel是第一类对象，可传参（内部实现为指针）或者作为结构成员
	req := NewRequest(10, 20, 30)
	Process(req)
	fmt.Println(<-req.ret)
	//说明:os.Args 返回命令⾏行参数，os.Exit 终⽌止进程。
	// 要获取正确的可执⾏行⽂文件路径，可⽤用 filepath.Abs(exec.LookPath(os.Args[0]))
	fmt.Println(os.Args)
}
//输出结果类似
//0 642081000
//1 642121000
//1 647307000
//0 647429000
//0 651648000
//1 651674000
///////////////
//60 （10+20+30）
