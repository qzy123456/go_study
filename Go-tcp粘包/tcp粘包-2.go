package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := encoder("s")
	fmt.Println(string(data))
	decoder(data)
}

/**
协议结构
4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。
*/
func decoder(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data len < 16")
		return
	}
	packetLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("包长度: %+v\n", packetLen)

	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("头长度: %+v\n", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("版本: %+v\n", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("业务操作码: %+v\n", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("序列号: %+v\n", sequence)

	body := string(data[16:])
	fmt.Printf("包内容: %+v\n", body)
}

func encoder(body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	response := make([]byte, packetLen)

	binary.BigEndian.PutUint32(response[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(response[4:6], uint16(headerLen))

	version := 5
	binary.BigEndian.PutUint16(response[6:8], uint16(version))
	operation := 6
	binary.BigEndian.PutUint32(response[8:12], uint32(operation))
	sequence := 7
	binary.BigEndian.PutUint32(response[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(response[16:], byteBody)

	return response
}