package main

import "fmt"

//实心
func shiXin(nums int){
	for i := 1; i <= nums; i++ {
		//打印几个空格
		for k := 1; k <= nums-i; k++ {
			fmt.Print(" ") //4 3 2 1 0
		}
		//打印多少个🌟
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*") //1 2 3 4 5
		}
		fmt.Println()
	}
}

//空心
func kongXin(nums int)  {
	//i 行数
	for i := 1; i <= nums; i++ {
		//打印多少空格
		for k := 1; k <= nums - i; k++{
			fmt.Print(" ")
		}
		//j 星的数量
		for j := 1; j <= 2 * i - 1 ; j++ {
			//第一个 最后一个 打印星星
			if j == 1 || j == 2 * i - 1 || i == nums {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main() {
    shiXin(5)
    kongXin(5)
}