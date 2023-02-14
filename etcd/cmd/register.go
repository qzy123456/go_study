package main

import (
	"flag"
	"github.com/zituocn/srd"
	"log"
)

func main() {
	flag.Parse()
	_, err := srd.NewRegister(&srd.RegisterOption{
		EtcdEndpoints: []string{"192.168.16.51:2379"},
		Lease:         3000,
		Schema:        "gk100-cache",
		Key:           "pod-ip",
		Val:           "10.10.10.2:10003",
	})
	if err != nil {
		log.Fatal(err)
	}
	select {}
}


