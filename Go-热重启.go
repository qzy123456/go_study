package main

import (
	"flag"
	"fmt"
	"github.com/fevin/gracehttp"
	"net/http"
)

func main() {
	//重启：kill -HUP pid
	//退出：kill -QUIT pid

	flag.Parse()
	sc := &Controller{}
	srv1 := &http.Server{
		Addr:    ":9094",
		Handler: sc,
	}
	gracehttp.AddServer(srv1, false, "", "")
	srv2 := &http.Server{
		Addr:    ":9093",
		Handler: sc,
	}
	gracehttp.AddServer(srv2, false, "", "")

	gracehttp.SetMaxConcurrentForOneServer(1)

	gracehttp.Run()

	fmt.Print("main over")
}

type Controller struct {
}

func (this *Controller) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello01"))
}