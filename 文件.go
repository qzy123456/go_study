package main

import (
	"fmt"
	"os"
)

func main() {
	//fout, err := os.Create("./xxx.txt") //新建文件
	fout, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fout.Close() //main函数结束前， 关闭文件

	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\n", "Hello go", i)
		fout.WriteString(outstr)     //写入string信息到文件
		fout.Write([]byte("abcd\n")) //写入byte类型的信息到文件
	}
	fin, err := os.Open("./xxx.txt") //打开文件
	if err != nil {
		fmt.Println(err)
	}
	defer fin.Close()

	buf := make([]byte, 1024) //开辟1024个字节的slice作为缓冲
	for {
		n, _ := fin.Read(buf) //读文件
		if n == 0 { //0表示已经到文件结束
			break
		}

		fmt.Println(string(buf)) //输出读取的内容
	}
}