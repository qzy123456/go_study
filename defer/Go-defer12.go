package main
import (
	"fmt"
)

func main()  {

	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}
//defer panic
//分析
//
//panic仅有最后一个可以被revover捕获。
//
//触发panic("panic")后defer顺序出栈执行，第一个被执行的defer中 会有panic("defer panic")异常语句，
//这个异常将会覆盖掉main中的异常panic("panic")，最后这个异常被第二个执行的defer捕获到。