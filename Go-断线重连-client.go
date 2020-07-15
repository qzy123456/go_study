package main
import (
	"net"
	"fmt"
	"bufio"
	"time"
)

func doTask(conn net.Conn) {
	for {
		fmt.Fprintf(conn,"test msg\n")
		msg,err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("recv data error")
			break
		}else{
			fmt.Println("recv msg : ",msg)
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {
	hostInfo := "127.0.0.1:12345"

	for {
		conn,err := net.Dial("tcp",hostInfo)
		fmt.Print("connect (",hostInfo)
		if err != nil {
			fmt.Println(") fail")
		}else{
			fmt.Println(") ok")
			defer conn.Close()
			doTask(conn)
		}
		time.Sleep(3 * time.Second)
	}
}