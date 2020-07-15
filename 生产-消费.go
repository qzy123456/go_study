package main
import (
	"fmt"
	"sync"
)
//实现一个生产者和消费者
/*生产者产生数据添加到通道里面，消费者消费数据从通道里面
不带缓存实现
*/
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go producers(&wg, ch)
	go consumers(&wg, ch)
	wg.Wait()
	fmt.Println("main exit")
	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
}
//生产者
func producers(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send:", i)
		ch <- i
	}
	close(ch)
	wg.Done()
}
//消费者
func consumers(wg *sync.WaitGroup, ch chan int) {
	for v := range ch {
		fmt.Println("recv:", v)
	}
	wg.Done()
}

type Command struct {
	Name    string    // 指令名称
	Var     *int      // 指令绑定的变量
	Comment string    // 指令的注释
}
