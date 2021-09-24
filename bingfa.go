package main

import (
	"fmt"
	"hash/crc32"
	"runtime"
	"time"
)

func says(s string) {
	for i := 0; i < 5; i++ {

		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go says("world")
	runtime.Gosched()
	//says("hello")
	fmt.Println("hello")
	//hash--crc32
	crc := crc32.ChecksumIEEE([]byte("12144143fgdfsdafsdf"))
	fmt.Println(crc)
	fmt.Println(int(crc))
	id := int(crc) % 4
	fmt.Print(id)

	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	//等待输入字符  才会退出
	var input string
	fmt.Scanln(&input)
}
