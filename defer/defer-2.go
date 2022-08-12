package main

func run() (string) {
	name := "Paul"
	defer sayHello(&name)
	name = "John"
	return name
}

func sayHello(name *string) {
	*name = "George"
	println("Hello " + *name)
}
func main() {
	name := run()
	println("return: name = " + name)
}
//Hello George
//return: name = John