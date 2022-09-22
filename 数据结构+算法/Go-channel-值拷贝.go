package main
import(
	"fmt"
	"time"
)
type user struct {
	name string
	age int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}
//Channel 发送和接收元素的本质是什么？
//All transfer of value on the go channels happens with the copy of value.
//就是说 channel 的发送和接收操作本质上都是 “值的拷贝”，无论是从 sender goroutine 的栈到 chan buf，
//还是从 chan buf 到 receiver goroutine，或者是直接从 sender goroutine 到 receiver goroutine。
//main 程序里，先把 g 发送到 c，根据 copy value 的本质，进入到 chan buf 里的就是 0x56420，
// 它是指针 g 的值（不是它指向的内容），所以打印从 channel 接收到的元素时，它就是 &{Ankur 25}。
// 因此，这里并不是将指针 g “发送” 到了 channel 里，只是拷贝它的值而已。
func main() {
	c := make(chan *user, 1)
	c <- g
	fmt.Println("先执行",g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	go printUser(c)
	go modifyUser(g)
	time.Sleep(3 * time.Second)
	fmt.Println(g)
	//先执行 &{Ankur 25}
	//modifyUser Received Vaule &{Ankur Anand 100}
	//printUser goRoutine called &{Ankur 25}
	//&{Anand 100}
}