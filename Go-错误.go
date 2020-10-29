package main
import ("fmt")

func fs() {
	fmt.Println("内层开始")               //第3步
	defer func() {
		fmt.Println("内层recover前的defer")     //第8步
	}()

	defer func() {
		fmt.Println("内层准备recover")           //第5步
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "内层", err) //第6步  这里err就是panic传入的内容（"异常信息"）
		}

		fmt.Println("内层完成recover")                //第7步
	}()

	defer func() {
		fmt.Println("内层异常前recover后的defer")  //第4步
	}()

	panic("异常信息")
    //下面其实都走不到
	defer func() {
		fmt.Println("内层异常后的defer")
	}()

	fmt.Println("内层异常后语句") //recover捕获的一级或者完全不捕获这里开始下面代码不会再执行
}

func main() {
	fmt.Println("外层开始")         //第1步
	defer func() {
		fmt.Println("外层准备recover")    //第11步
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "外层", err) // err已经在上一级的函数中捕获了，这里没有异常，只是例行先执行defer，然后执行后面的代码
		} else {
			fmt.Println("外层没做啥事")    //第12步
		}
		fmt.Println("外层完成recover")    //第13步
	}()
	fmt.Println("外层即将异常")    //第2步
	fs()
	fmt.Println("外层异常后")      //第9步
	defer func() {
		fmt.Println("外层异常后defer")  //第10步
	}()
}
//外层开始
//外层即将异常
//内层开始
//内层异常前recover后的defer
//内层准备recover
//"内层"-"异常信息"
//内层完成recover
//内层recover前的defer
//外层异常后
//外层异常后defer
//外层准备recover
//外层没做啥事
//外层完成recover
