package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var (
	cache = make([]byte, 0) // 缓存
)

const (
	HeaderLen = 4 // 每个数据包前4个字节保存一个整形，该整形记录数据包的字节长度
)

// 打包
func Packet(buf []byte) []byte {
	return append(intToBytes(len(buf)), buf...)
}

// 解包
func Unpack(buf []byte) []byte {
	buf = append(cache, buf...)
	length := len(buf)

	messageLength := bytesToInt(buf[:HeaderLen])
	// 当前数据包的总长度
	total := HeaderLen + messageLength
	if length < total {
		cache = buf
		return []byte{}
	} else if length == total {
		cache = []byte{}
		return buf[HeaderLen:]
	} else {
		cache = buf[total:]
		return buf[HeaderLen:total]
	}
}

//整形转换成字节
func intToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func main() {
	var buf []byte
	// 打包一个消息
	for i:=0;i<100 ;i++  {
		s := fmt.Sprintf("%s%d","this is message ",i)
		buf = Packet([]byte(s))
	}
	fmt.Println(string(buf))
	// 解包一个消息
	buf = Unpack(buf)
	fmt.Println(string(buf))

}
