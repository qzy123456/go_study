package main

func run() (x string) {
	name := "Paul"
	x = name
	aFun := func(x string) {
		println("Hello " + x)
		x = "George"
		println("Hello " + x)
	}
	defer aFun(x)
	name = "John"
	return name
}
func main() {
	name := run()
	println("return: name = " + name)
}

//Hello Paul
//Hello George
//return: name = John
