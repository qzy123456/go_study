package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)
type student struct {
	Name string
	Age uint8
	Address string
}
//准备编码的数据
type P struct {
	    X, Y, Z int
	    Name    string
	}

//接收解码结果的结构
type Q struct {
	    X, Y *int32
	    Name string
	}

func main() {
	    //初始化一个数据
	data := P{3, 4, 5, "CloudGeek"}
	//编码后得到buf字节切片
	buf := encode(data)
	//用于接收解码数据
	var q *Q
	//解码操作
	q = decode(buf)
	//"CloudGeek": {3,4}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
	/////////////////第二种写法//////////////////////////
	//序列化
	s1 := student{"张三", 18, "江苏省"}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) //创建编码器
	err1 := encoder.Encode(&s1)        //编码
	if err1 != nil {
		log.Panic(err1)
	}
	fmt.Printf("序列化后：%x\n", buffer.Bytes())

	//反序列化
	byteEn := buffer.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(byteEn)) //创建解密器
	var s2 student
	err2 := decoder.Decode(&s2) //解密
	if err2 != nil {
		log.Panic(err2)
	}
	fmt.Println("反序列化后：", s2)

}

func encode(data interface{}) *bytes.Buffer {
	    //Buffer类型实现了io.Writer接口
	    var buf bytes.Buffer
	    //得到编码器
	    enc := gob.NewEncoder(&buf)
	    //调用编码器的Encode方法来编码数据data
	    enc.Encode(data)
	    //编码后的结果放在buf中
	    return &buf
	}

func decode(data interface{}) *Q {
	    d := data.(io.Reader)
	    //获取一个解码器，参数需要实现io.Reader接口
	    dec := gob.NewDecoder(d)
	    var q Q
	    //调用解码器的Decode方法将数据解码，用Q类型的q来接收
	    dec.Decode(&q)
	    return &q
	}