package main

import "fmt"

func main() {
	ff()
	//这一句肯定是最后执行，因为最后才到main函数
	fmt.Println("Returned normally from f.") //12
}

func ff() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r) // 11
		}
	}()
	fmt.Println("Calling g.") // 1
	gg(0)
	fmt.Println("Returned normally from g.") //这句其实永远不会执行，出错之后，就走不到下面
}

func gg(i int) {
	if i > 3 {
		fmt.Println("Panicking!") //6
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i) //7，8，9，10
	fmt.Println("Printing in g", i) //2 ，3，4 ，5
	gg(i + 1)
}
//Calling g.
//Printing in g 0
//Printing in g 1
//Printing in g 2
//Printing in g 3
//Panicking!
//Defer in g 3
//Defer in g 2
//Defer in g 1
//Defer in g 0
//Recovered in f 4
//Returned normally from f.