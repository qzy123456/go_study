// main.go

package main

import "wire/service"

func main() {
	client := service.BuildClient()
	client.Foo.Foo()
	client.Bar.Bar()
}