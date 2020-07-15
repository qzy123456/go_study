package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"os"
	"log"
	"../go/pb/pb"
)

func write() {
	c1 := &pb.Class{
		Num: 1,
		Students: []*pb.Student{
			{Name: "xiaoming", Age: 21, Sex: pb.Sex_MAN},
			{Name: "xiaohua", Age: 21, Sex: pb.Sex_WOMAN},
			{Name: "xiaojin", Age: 21, Sex: pb.Sex_MAN},
		},
	}

	// 使用protobuf工具把struct数据类型格式化成字节数组（压缩和编码）
	data, _ := proto.Marshal(c1)

	// 把字节数组写入到文件中
	ioutil.WriteFile("test.txt", data, os.ModePerm)
}

func read1() {
	// 以字节数组的形式读取文件内容
	data, _ := ioutil.ReadFile("test.txt")

	class := new(pb.Class)

	// 使用protobuf工具把字节数组解码成struct(解码)
	if err := proto.Unmarshal(data, class); err != nil{
		fmt.Println("un json 出错了")
	}

	log.Println(class.Num)
	for _, v := range class.Students {
		log.Println(v.Name, v.Age, v.Sex)
	}
}

func main() {
	write()
	read1()
}

