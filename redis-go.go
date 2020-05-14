package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(1),//手动选择数据库，不选默认就是0
	)
	if err != nil {
		fmt.Println("Connect to redis failed ,cause by >>>", err)
		return
	}
	defer conn.Close()

	//如果db有密码，可以设置
	//if _,err := conn.Do("AUTH","password");err !=nil{
	//	fmt.Println("connect db by pwd failed >>>",err)
	//}

	//写入值{"test-Key":"100"}
	_, err = conn.Do("SET", "test-Key", 100, "EX", "500")
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}

	time.Sleep(10 * time.Second)
	//检查是否存在key值
	exists, err := redis.Bool(conn.Do("EXISTS", "test-Key"))
	if err != nil {
		fmt.Println("illegal exception")
	}
	fmt.Printf("exists or not: %v \n", exists)

	//read value
	v, err := redis.String(conn.Do("GET", "test-Key"))
	if err != nil {
		fmt.Println("redis get value failed >>>", err)
	}
	fmt.Println("get value: ", v)

	//del kv
	_, err = conn.Do("DEL", "test-Key")
	if err != nil {
		fmt.Println("redis delelte value failed >>>", err)
	}
}
