package main

import "fmt"

//遇到左括号就进栈push，遇到右括号并且栈顶为与之对应的左括号，就把栈顶元素出栈。
// 最后看栈里面还有没有其他元素，如果为空，即匹配。
func isValids(s string) bool {
	//空字符串为true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, value := range s {
		if (value == '[') || (value == '(') || (value == '{') {
			stack = append(stack, value)
		} else if (value == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(value == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(value == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s1 := "([)]"
	s2 := "{[]}"
	fmt.Println(isValids(s1))
	fmt.Println(isValids(s2))
}
