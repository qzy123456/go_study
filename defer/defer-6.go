package main

func run() (x string) {
	name := "Paul"
	aFun := func() {
		println("Hello " + x)
		x = "George"
		println("Hello " + x)
	}
	defer aFun()
	name = "John"
	return name
}
func main() {
	name := run()
	println("return: name = " + name)
}

//Hello John
//Hello George
//return: name = George
