package main
import (
	"errors"
	"fmt")
type People11 interface {
	Speak(string) string
}
type Stduent struct{

}
func(stu *Stduent) Speak(think string)(talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	//var peo People11 = Stduent{}  //两个类型不一致   会报错的，直接编译的时候 就不会通过
	var peo  = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	//BBBBBBB
	fmt.Println(Test1())  //1
	fmt.Println(Test2())  //2
	fmt.Println(Test3())  //4

	e1()
	e2()
	e3()
   //<nil>
   //e2 defer err
   //<nil>
}

type People22 interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People22 {
	var stu *Student
	return stu
}

func Test1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}
func Test2() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 2
}
func Test3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}