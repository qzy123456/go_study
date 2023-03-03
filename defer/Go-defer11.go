package main
import "fmt"

func main() {
	a()
}

func a(){
	defer a1()
	defer a2()
	panic("a")
}
func a1(){
	fmt.Println("a1")
}

func a2(){
	defer b1()
	panic("a2")
}

func b1(){
	p:=recover()
	fmt.Println(p)

}

//a2
//a1
//panic: a