package main

import (
	"fmt"
	"errors"
	"strconv"
)

//用栈计算算数算数表达式的结果

type Stack struct {
	MaxSize int
	Top int
	Array [20]int
}

//入栈
func (this *Stack) Push(val int) error {
	if this.IsFull() {
		return errors.New("该栈已满！")
	}
	this.Top++
	this.Array[this.Top] = val
	return nil
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("该栈已空！")
	}
	val = this.Array[this.Top]
	this.Top--
	return val, nil
}

//判断栈已满
func (this *Stack) IsFull() bool {
	return this.Top == this.MaxSize - 1
}

//判断该栈为空
func (this *Stack) IsEmpty() bool {
	return this.Top == -1
}

//显示栈里的值
func (this *Stack) List() {
	if this.IsEmpty() {
		fmt.Println("该栈为空！")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("%d\t",this.Array[i])
	}
}


//判断如果不是运算符就是数字
//如果是运算符返回true，是数字返回false
//*，+，-，/ 运算发依次对应的ASCII码为42，43，45，47
func (this *Stack) IsOper(val int) bool {
	return val == 42 || val == 43 || val == 45 || val == 47
}

//判断运算符的优先级
func (this *Stack) Priority(val int) int {
	if val == 23 || val == 45 {
		return 0
	} else if val == 42 || val == 47 {
		return 1
	} else {
		return -1
	}
}

//将从栈里弹出的两个数和运算符进行计算
func (this *Stack) Cal(a, b , oper int) int {
	switch oper {
	case 42:
		return b * a
	case 43:
		return b + a
	case 45:
		return b - a
	case 47:
		return b / a
	default:
		return 0
	}
}


func main() {
	//数字的栈
	number := &Stack {
		MaxSize : 20,
		Top : -1,
	}
	//运算符的栈
	oper := &Stack {
		MaxSize : 20,
		Top : -1,
	}

	fmt.Println("请输入你要计算的算数表达式（仅支持+，-，*，/）：")
	var expr string
	fmt.Scan(&expr)
	//字符串expr的下标
	index := 0
	operTemp := 0
	keepNum := ""
	for {
		str := expr[index:index+1]

		operTemp = int([]byte(str)[0])
		//判断是运算符还是数字
		if oper.IsOper(operTemp) {
			//如果运算符栈里没有值，就直接入栈
			if oper.Top == -1 {
				oper.Push(operTemp)
			} else {
				//如果栈里已有运算符则判断两个运算符的优先级
				if oper.Priority(oper.Array[oper.Top]) >= oper.Priority(operTemp) {
					//从数字栈弹出两个数
					a,_ := number.Pop()
					b,_ := number.Pop()
					//从运算符栈弹出一个运算符
					c,_ := oper.Pop()
					//将运算的结果重新入数字栈
					number.Push(number.Cal(a, b, c))
					//将运算符入栈
					oper.Push(operTemp)
				} else {
					oper.Push(operTemp)
				}
			}
		} else {
			keepNum += str
			//如果index==len(expr)则表示expr已读到最后，直接入数字栈
			if index == len(expr) - 1 {
				data,_ := strconv.ParseInt(keepNum,10, 64)
				number.Push(int(data))
			} else {
				//判断str后一位是不是运算符，如果是就直接入数字栈，如果不是keepNum就拼接str
				if oper.IsOper(int([]byte(expr[index+1:index+2])[0])) {
					data,_ := strconv.ParseInt(keepNum,10, 64)
					number.Push(int(data))
					keepNum = ""
				}
			}
		}
		index++
		if index == len(expr) {
			break
		}
	}

	//将数字栈里的数依次弹出两个，运算符栈弹出一个 依次计算
	for {
		//运算符栈为空时推出
		if oper.Top == -1 {
			break
		}
		a,_ := number.Pop()
		b,_ := number.Pop()
		c,_ := oper.Pop()
		number.Push(number.Cal(a, b, c))
	}
	res,_ := number.Pop()
	fmt.Printf("%s=%d",expr,res)
}