package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"
)
const HEADER = "BEGIN"
// 每个消息(包括头部)的最大长度， 这里最大可以设置4G
const HeadSize = 4

type Buffer struct {
	reader    io.Reader
	header    string
	buf       []byte
	bufLength int
	start     int
	end       int
}

func NewBuffer(reader io.Reader, header string, len int) *Buffer {
	buf := make([]byte, len)
	return &Buffer{reader, header, buf, len, 0, 0}
}

// grow 将有用的字节前移
func (buffer *Buffer) grow() {
	if buffer.start == 0 {
		return
	}
	copy(buffer.buf, buffer.buf[buffer.start:buffer.end])
	buffer.end -= buffer.start
	buffer.start = 0
}

func (buffer *Buffer) len() int {
	return buffer.end - buffer.start
}

//返回n个字节，而不产生移位
func (buffer *Buffer) seek(n int) ([]byte, error) {
	if buffer.end-buffer.start >= n {
		buf := buffer.buf[buffer.start : buffer.start+n]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

//舍弃offset个字段，读取n个字段
func (buffer *Buffer) read(offset, n int) []byte {
	buffer.start += offset
	buf := buffer.buf[buffer.start : buffer.start+n]
	buffer.start += n
	return buf
}

//从reader里面读取数据，如果reader阻塞，会发生阻塞
func (buffer *Buffer) readFromReader() error {
	if buffer.end == buffer.bufLength {
		return errors.New(fmt.Sprintf("一个完整的数据包太长已经超过你定义的example.BUFFER_LENGTH(%d)\n", buffer.bufLength))
	}
	n, err := buffer.reader.Read(buffer.buf[buffer.end:])
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second) // 便于观察这里sleep了一下
	buffer.end += n
	return nil
}

func (buffer *Buffer) Read(msg chan string) error {
	for {
		buffer.grow()                   // 移动数据
		err1 := buffer.readFromReader() // 读数据拼接到定额缓存后面
		if err1 != nil {
			return err1
		}
		// 检查定额缓存里面的数据有几个消息(可能不到1个，可能连一个消息头都不够，可能有几个完整消息+一个消息的部分)
		err2 := buffer.checkMsg(msg)
		if err2 != nil {
			return err2
		}
	}
}

func (buffer *Buffer) checkMsg(msg chan string) error {
	HeaderLen := HeadSize + len(buffer.header)
	headBuf, err1 := buffer.seek(HeaderLen)
	if err1 != nil { // 一个消息头都不够， 跳出去继续读吧, 但是这不是一种错误
		return nil
	}
	if string(headBuf[:len(buffer.header)]) == buffer.header { // 判断消息头正确性

	} else {
		return errors.New("消息头部不正确")
	}
	contentSize := int(binary.BigEndian.Uint32(headBuf[len(buffer.header):]))
	if buffer.len() >= contentSize-HeaderLen { // 一个消息体也是够的
		contentBuf := buffer.read(HeaderLen, contentSize) // 把消息读出来，把start往后移
		msg <- string(contentBuf)
		// 递归，看剩下的还够一个消息不
		err3 := buffer.checkMsg(msg)
		if err3 != nil {
			return err3
		}
	}
	return nil
}
