package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	 //"log"
)

type handle struct{
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
func startServer() { //被代理的服务器host和port
	h := &handle{host: "127.0.0.1", port: "80"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Info("ListenAndServe: ", err)
	}
}
func main() {
	startServer()
}
