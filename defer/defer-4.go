package main

func run() (x string) {
	name := "Paul"
	x = name
	defer sayHello(x)
	name = "John"
	return name
}

func sayHello(name string) {
	println("Hello " + name)
	name = "George"
	println("Hello " + name)
}
func main() {
	name := run()
	println("return: name = " + name)
}

//Hello Paul
//Hello George
//return: name = John
