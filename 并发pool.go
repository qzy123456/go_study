package main
import (
	"flag"
	"fmt"
	"net/url"
	"runtime"

	"sync"
	"strconv"
	// "time"

	"github.com/gorilla/websocket"
)
type Pool struct {
	Queue chan int
	Wg    *sync.WaitGroup
}
var (
	addr = flag.String("addr", "192.168.31.68:8080", "http service address")
	n    = flag.String("n", "1", "请求总数")
	c    = flag.String("c", "1", "一次请求并发数")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/goldminer"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	n_int, _ := strconv.Atoi(*n)
	c_int, _ := strconv.Atoi(*c)
	fmt.Println("n:", n_int)
	fmt.Println("c:", c_int)
	pool := NewPool(c_int, n_int)
	for i := 0; i < n_int; i++ {
		go send(conn, pool)
	}
	pool.Wg.Wait()
}

func send(conn *websocket.Conn, pool *Pool) {
	defer func() {
		pool.DelOne()
	}()
	pool.AddOne()
	content := `{"command": "login","token": "","data": {"username": "gc","password": "123456","code": "482E74902DFF0A84538A6676BB866A66","openid": "77EAAEADA573A37A394379FFC58C67F6","bind": 2}}`
	err := conn.WriteMessage(websocket.TextMessage, []byte(content))
	if err != nil {
		fmt.Println("write err:", err)
	}

}


// 创建并发控制池, 设置并发数量与总数量
func NewPool(cap, total int) *Pool {
	if cap < 1 {
		cap = 1
	}
	p := &Pool{
		Queue: make(chan int, cap),
		Wg:    new(sync.WaitGroup),
	}
	p.Wg.Add(total)
	return p
}

// 向并发队列中添加一个
func (p *Pool) AddOne() {
	p.Queue <- 1
}

// 并发队列中释放一个, 并从总数量中减去一个
func (p *Pool) DelOne() {
	<-p.Queue
	p.Wg.Done()
}



