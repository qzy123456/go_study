
package main

func run() {
	name := "Paul"
	defer sayHello(name)
	name = "John"
}

func sayHello(name string) {
	println("Hello " + name)
}

func main() {
	run()
}
//Hello Paul