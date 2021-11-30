package main

import "fmt"
//   如果是 ‘.’, 则需要 ‘.‘没有出现过，并且 ‘e/E’ 没有出现过，才会进行标记
//   如果是 ‘e/E’, 则需要 ‘e/E’没有出现过，并且前面出现过数字，才会进行标记
//      如果是 ‘+/-’, 则需要是第一个字符，或者前一个字符是 ‘e/E’，才会进行标记，并重置数字出现的标识
//       最后返回时，需要字符串中至少出现过数字，避免下列case: s == ‘.’ or ‘e/E’ or ‘+/e’ and etc…

func isNumber(s string) bool {
	//标记是不是数字 小数点 以及e E
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false // reJudge integer after 'e' or 'E'
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	// avoid case: s == '.' or 'e/E' or '+/-' and etc...
	// string s must have num
	return numFlag
}
func main() {
	fmt.Println(isNumber("e"))
}
