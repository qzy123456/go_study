package main

import (
	"encoding/gob"
	"fmt"
	"strings"
	"bytes"
)

func convert(){
	stringSlice := []string{"通知中心","perfect!"}

	stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null

	fmt.Println([]byte(stringByte))

	fmt.Println(string([]byte(stringByte)))
}

func convert2(){
	stringSlice := []string{"通知中心","perfect!"}

	buffer := &bytes.Buffer{}

	gob.NewEncoder(buffer).Encode(stringSlice)
	byteSlice := buffer.Bytes()
	fmt.Printf("%q\n", byteSlice)

	fmt.Println("---------------------------")

	backToStringSlice := []string{}
	gob.NewDecoder(buffer).Decode(&backToStringSlice)
	fmt.Printf("%v\n", backToStringSlice)
}
func main() {
	convert()
	convert2()
}