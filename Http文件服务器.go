package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	address := GetIntranetIp()
	fmt.Println("本机ip地址列表：")
	for _, item := range address {
		fmt.Println(item)
	}
	http.Handle("/", http.FileServer(http.Dir("./")))
	fmt.Printf("文件共享服务开启，监听8888端口\n请使用浏览器打开:http://ip地址:8888,eg:http://%s:8888\n请不要关闭此程序，祝使用愉快\n", address[0])
	go func() {
		time.Sleep(2000)
		loclstr := fmt.Sprintf("http://%s:8888", address[0])
		cmd := exec.Command("cmd", "/C", "start "+loclstr)
		cmd.Run()
	}()
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("err:", err)
	}
}

func GetIntranetIp() (r []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				r = append(r, ipnet.IP.String())
			}
		}
	}
	return
}