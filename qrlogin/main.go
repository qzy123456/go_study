package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	config      *Config
	channelPool *ChannelPool
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var configFile string
	root := filepath.Dir(os.Args[0])
	flag.StringVar(&configFile, "config", filepath.Join(root, "config.json"), "configuration file path, default is ./config.json")
	flag.Parse()

	config = NewConfigFromFile(configFile)
	channelPool = NewChannelPool()

	http.Handle("/qrlogin/public/", http.StripPrefix("/qrlogin/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/qrlogin", IndexHandler)
	http.HandleFunc("/qrlogin/get_channel", GetChannel)
	http.HandleFunc("/qrlogin/check", CheckLogin)
	http.HandleFunc("/qrlogin/login", ReadyToLogin)
	http.HandleFunc("/qrlogin/confirm", ConfirmLogin)

	log.Println("ListenAndServe: ", config.Listen)
	err := http.ListenAndServe(config.Listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
