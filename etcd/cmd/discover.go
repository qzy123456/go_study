package main

import (
	"github.com/zituocn/srd"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dis, err := srd.NewDiscovery("gk100-cache", "pod-ip", []string{"192.168.16.51:2379"}, 3)
	if err != nil {
		log.Fatal(err)
	}
	err = dis.Builder()
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	for {
		select {
		case <-time.Tick(3 * time.Second):
			log.Println(dis.GetValues())
		case <-c:
			log.Println("server discovery exit")
			return
		}
	}
}
