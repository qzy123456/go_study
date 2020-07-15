package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"os/signal"
)

var chs = make(chan string)
var cin = bufio.NewScanner(os.Stdin)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			os.Exit(0)

		}
	}()
	go clientWriter(conn, chs)
	for cin.Scan() {
		op := cin.Text()
		op = strings.Replace(op, " ", "", -1)
		switch op {
		case "register":
			register()
		case "login":
			login()
		case "logoff":
			chs <- "logoff"
		case "exit":
			os.Exit(0)
		case "add":
			adds()
		case "delete":
			delete()
		case "list":
			chs <- "list"
		case "sendTo":
			sendTo()
		case "sendAll":
			sendAll()
		default:
			chs <- "none"
		}

	}
	<-done // wait for background goroutine to finish
}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
func register() {
	chs <- "register"
	var userName, password, tel string
	fmt.Print("请输入用户名：")
	if cin.Scan() {
		userName = cin.Text()
	}
	fmt.Print("请输入密码：")
	if cin.Scan() {
		password = cin.Text()
	}
	fmt.Print("请输入手机号：")
	if cin.Scan() {
		tel = cin.Text()
	}
	chs <- userName
	chs <- password
	chs <- tel
}

func login() {
	chs <- "login"
	var userName, password string
	fmt.Print("请输入用户名：")
	if cin.Scan() {
		userName = cin.Text()
	}
	fmt.Print("请输入密码：")
	if cin.Scan() {
		password = cin.Text()
	}
	chs <- userName
	chs <- password
}
func adds() {
	var userName string
	fmt.Print("请输入要添加好友的用户名：")
	if cin.Scan() {
		userName = cin.Text()
	}
	chs <- "add"
	chs <- userName
}
func delete() {
	fmt.Print("请输入要删除好友的用户名：")
	var userName string
	if cin.Scan() {
		userName = cin.Text()
	}
	chs <- "delete"
	chs <- userName
}
func sendTo() {
	var userName, content string
	fmt.Print("请输入好友用户名：")
	if cin.Scan() {
		userName = cin.Text()
	}
	fmt.Print("请输入消息：")
	if cin.Scan() {
		content = cin.Text()
	}
	chs <- "sendTo"
	chs <- userName
	chs <- content
}
func sendAll() {
	var content string
	fmt.Print("请输入消息：")
	if cin.Scan() {
		content = cin.Text()
	}
	chs <- "sendAll"
	chs <- content

}
