package main

import "fmt"

func surroundingFuncEvaluatedNotInvoked(init int) int {
	fmt.Printf("1.init=%d\n",init)

	defer func() {
		fmt.Printf("2.init=%d\n",init)

		init ++

		fmt.Printf("3.init=%d\n",init)
	}()

	fmt.Printf("4.init=%d\n",init)

	return init
}

func noDeferFuncOrderWhenReturn() (result int) {
	func() {
		// 1. before : result = 0
		fmt.Printf("before : result = %v\n", result)

		result++

		// 2. after : result = 1
		fmt.Printf("after : result = %v\n", result)
	}()

	// 3. return : result = 1
	fmt.Printf("return : result = %v\n", result)

	return 0
}


func deferFuncWithAnonymousReturnValue() int {
	var retVal int
	defer func() {
		retVal++
	}()
	return 0
}

func deferFuncWithNamedReturnValue() (retVal int) {
	defer func() {
		retVal++
	}()
	return 0
}
//go 语言官方示例
func fibonacci1() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main()  {
	surroundingFuncEvaluatedNotInvoked(1)
	noDeferFuncOrderWhenReturn()
	fmt.Println(deferFuncWithAnonymousReturnValue()) //0
	fmt.Println(deferFuncWithNamedReturnValue())   // 1
  	//1.init=1
	//4.init=1
	//2.init=1
	//3.init=2
	//before : result = 0
	//after : result = 1
	//return : result = 1
	//0
	//1
	f := fibonacci1()
	for i := 0; i < 20; i++ {
		fmt.Print(f(), " ")
	}
}