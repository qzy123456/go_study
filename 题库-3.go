package main
import (   "fmt")
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
