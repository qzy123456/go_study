package main

import "fmt"

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res   = []string{}
	final = 0
)

// 解法一 DFS
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
	return
}

// 解法二 非递归
func letterCombinations_(digits string) []string {
	if digits == "" {
		return []string{}
	}
	index := digits[0] - '0'
	letter := letterMap[index]
	tmp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res, "")
		}
		for j := 0; j < len(res); j++ {
			tmp = append(tmp, res[j]+string(letter[i]))
		}
	}
	res = tmp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		tmp = res
		res = []string{}
	}
	return tmp
}

// 解法三 回溯（参考回溯模板，类似DFS）
var result []string
var dict = map[string][]string{
	"2" : []string{"a","b","c"},
	"3" : []string{"d", "e", "f"},
	"4" : []string{"g", "h", "i"},
	"5" : []string{"j", "k", "l"},
	"6" : []string{"m", "n", "o"},
	"7" : []string{"p", "q", "r", "s"},
	"8" : []string{"t", "u", "v"},
	"9" : []string{"w", "x", "y", "z"},
}

func letterCombinationsBT(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	letterFunc("", digits)
	return result
}

func letterFunc(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}

	k := digits[0:1] //第一个按键
	digits = digits[1:] //后面n个按键
	for i := 0; i < len(dict[k]); i++ { //第一个按键上的n个字母
		res += dict[k][i]
		letterFunc(res, digits)
		res = res[0 : len(res)-1]   //最后一个是一个单独字母 要抛弃
	}
}

func main() {
	s:="23"
	fmt.Println(letterCombinations(s))
	fmt.Println(letterCombinations_(s))
	fmt.Println(letterCombinationsBT(s))
}