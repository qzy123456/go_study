package main

import (
	"bytes"
	"fmt"
	"bufio"
)

func main() {
	
}
////Read() 方法的功能是读取数据，并存放到字节切片 p 中。Read() 执行结束会返回已读取的字节数，
// 因为最多只调用底层的 io.Reader 一次，所以返回的 n 可能小于 len(p)，当字节流结束时，n 为 0，err 为 io. EOF。该方法原型如下：
func read()  {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
}
//ReadByte() 方法的功能是读取并返回一个字节，如果没有字节可读，则返回错误信息。该方法原型如下：
func readByte()  {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}
//ReadBytes() 方法的功能是读取数据直到遇到第一个分隔符“delim”，并返回读取的字节序列（包括“delim”）。
// 如果 ReadBytes 在读到第一个“delim”之前出错，它返回已读取的数据和那个错误（通常是 io.EOF）。
// 只有当返回的数据不以“delim”结尾时，返回的 err 才不为空值。该方法原型如下：
func readBytes()  {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)
}
//ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者 ReadString('\n')。
// ReadLine 返回一行，不包括结尾的回车字符，如果一行太长（超过缓冲区长度），
// 参数 isPrefix 会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回。
func readLine()  {
	data := []byte("Golang is a beautiful language. \r\n I like it!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)
}
//ReadRune() 方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数。
// 如果编码错误，ReadRune 只读取一个字节并返回 unicode.ReplacementChar(U+FFFD) 和长度1
func readRune()  {
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}
//ReadSlice() 方法的功能是读取数据直到分隔符“delim”处，并返回读取数据的字节切片，
// 下次读取数据时返回的切片会失效。如果 ReadSlice 在查找到“delim”之前遇到错误，它返回读取的所有数据和那个错误（通常是 io.EOF）。
//如果缓冲区满时也没有查找到“delim”，则返回 ErrBufferFull 错误。
// ReadSlice 返回的数据会在下次 I/O 操作时被覆盖，大多数调用者应该使用 ReadBytes 或者 ReadString。
// 只有当 line 不以“delim”结尾时，ReadSlice 才会返回非空 err。该方法原型如下
func readSlice()  {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}
//ReadString() 方法的功能是读取数据直到分隔符“delim”第一次出现，并返回一个包含“delim”的字符串。
// 如果 ReadString 在读取到“delim”前遇到错误，它返回已读字符串和那个错误（通常是 io.EOF）。
// 只有当返回的字符串不以“delim”结尾时，ReadString 才返回非空 err。该方法原型如下：
func readString()  {
	data := []byte("C语言中文网, Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println(line, err)
}
