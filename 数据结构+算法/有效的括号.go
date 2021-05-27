package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0{
		return true
	}
	//栈
	var stackList []string
	//配对字典
	judgeMap := map[string]string{")":"(", "}":"{", "]":"["}
	//把字符串的每个字符放进栈中，每放一个就判断与前一个是不是配对的
	for i:=0; i<len(s); i++{
		if len(stackList) == 0{
			stackList = append(stackList, string(s[i]))
			fmt.Println(stackList)
		} else {
			//判断是否配对
			//如果是相同的话，那就去除栈的最后一个元素
			//如果不相同的话，那就把源字符串的对应元素加进栈中
			//这里有个小技巧，就是每次我们放进容器的字符，当配对成功的时候，肯定是塞进右边的符号，
			//所以可以构造一个以右边括号为key，左边括号为值得字典
			if stackList[len(stackList)-1] == judgeMap[string(s[i])] {
				//fmt.Println(stackList,judgeMap[string(s[i])])
				stackList = stackList[:len(stackList)-1]
			} else {
				stackList = append(stackList, string(s[i]))
			}
		}
	}
	if len(stackList) != 0{
		return false
	} else {
		return true
	}
}

func main() {
	test := "()"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "()[]{}"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "(]"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "([)]"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "{[]}"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "{[qdag]}"
	fmt.Printf("judge is: %v\n", isValid(test))
}
