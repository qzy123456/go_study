package main

import "fmt"

func main() {
	//&符号的意思是对变量取地址，如：变量a的地址是&a 得到的是Ox0000000这种的数据
	//*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了
	//*和&可以互相抵消,同时注意，*&可以抵消掉，但&*是不可以抵消的
	//a和*&a是一样的，都是a的值，值为1 (因为*&互相抵消掉了)
	//同理，a和*&*&*&*&a是一样的，都是1 (因为4个*&互相抵消掉了)
	//因为有
	//var b *int = &a
	//所以
	//a和*&a和*b是一样的，都是a的值，值为1 (把b当做&a看)
	//因为有
	//var c **int = &b
	//所以
	//**c和**&b是一样的，把&约去后
	//会发现**c和*b是一样的 (从这里也不难看出，c和b也是一样的)
	//又因为上面得到的&a和b是一样的 所以**c和*&a是一样的，再次把*&约去后**c和a是一样的，都是1
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a) 						//1
	fmt.Println("&a = ",&a) 						//地址
	fmt.Println("*&a = ",*&a) 					//1
	fmt.Println("b = ",b) 						//地址
	fmt.Println("&b = ",&b) 						//1
	fmt.Println("*&b = ",*&b) 					//地址
	fmt.Println("*b = ",*b)						//1
	fmt.Println("c = ",c)  						//地址
	fmt.Println("*c = ",*c)						//地址
	fmt.Println("&c = ",&c) 						//地址
	fmt.Println("*&c = ",*&c)					//地址
	fmt.Println("**c = ",**c) 					//1
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c) //1
	fmt.Println("x = ",x)						//1

/**
	a 			=  1
	&a 			=  0xc00008e000
	*&a 		=  1
	b 			=  0xc00008e000
	&b 			=  0xc000088010
	*&b 		=  0xc00008e000
	*b 			=  1
	c 			=  0xc000088010
	*c		 	=  0xc00008e000
	&c 			=  0xc000088018
	*&c		 	=  0xc000088010
	**c 		=  1
	***&*&*&*&c =  1
	x 			=  1
*/
}