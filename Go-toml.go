package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type database struct {
	Server string
	Ports  []int
}

type config struct {
	Database database
}

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed,err:", err)
	}
	var c config
	viper.Unmarshal(&c)
	fmt.Println(c.Database.Ports[1])
}