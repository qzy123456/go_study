package main

func run() (string) {
	name := "Paul"
	aFun := func() {
		println("Hello " + name)
		name = "George"
		println("Hello " + name)
	}
	name = "John"
	aFun()
	return name
}
func main() {
	name := run()
	println("return: name = " + name)
}

//Hello John
//Hello George
//return: name = George
