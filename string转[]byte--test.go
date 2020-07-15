package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis"
)
func Redis() (*redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client;
}

func main()  {
	client := Redis()
	//获取所有的离线消息列表
	messagesList,err:=client.LRange("22go-chat",0,-1).Result()
	if err!=nil {
		return
	}


	cc,_ :=ComputeHash(messagesList)
	fmt.Println(cc)

}

func ComputeHash(key []string) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	data := buf.Bytes()
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&key)
	if err != nil {
		fmt.Println("decode error:", err)
	}
	fmt.Println("decode",key)
	return data, nil

}