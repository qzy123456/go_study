package main

import (
	"errors"
	"fmt"
)
func foo_1() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()

	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)

	err = errors.New("c")
	return err
}

func foo_2() error {
	var err error
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()

	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)

	err = errors.New("c")
	return err
}

func main() {
	fmt.Println(foo_1()) //nil c a
	//第一步：函数foo_1()执行到return关键字时，先准备好返回值，err=c
	//第二步：进入第二个defer函数，这里面进行了值拷贝，将err拷贝给了e，所以这里打印的e是nil,之后将e赋值也是不影响err的值的，err还是c
	//第三步：进入第一个defer函数，这里先打印了err，还是c，然后对err进行了赋值，err值变为a
	//第四步：经过defer的操作，现在的err值变为了a，故返回给调用者是a
	fmt.Println(foo_2()) //nil c c
	//这里的返回值是匿名返回写法，这种情况下，return会首先创建一个临时变量temp，然后将err赋值给temp,
	//后续两个defer中对err的操作并不会影响到temp中的值，所以这里返回给调用者的是c
}