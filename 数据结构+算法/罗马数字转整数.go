package main

import "fmt"

func romanToInt(s string) int {
	//建立map，方便查值
	strToNum := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	s = s + "I"
	//由左到右求和
	for i := 0; i < len(s) - 1; i++ {
		//遇到左边比右边小的要减去左边的
		if strToNum[s[i]] < strToNum[s[i + 1]] {
			ans += strToNum[s[i + 1]] - strToNum[s[i]]
			i++
			fmt.Println("1:", ans)
		} else {
			ans += strToNum[s[i]]
			fmt.Println("2:", ans)
		}

	}
	return ans
}

func romanToInt2(s string) int {
	specialRomanStringMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900} //特殊罗马数字
	romanStringMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}   //正常罗马数字
	result := 0
	for len(s) != 0 { //对字符串循环
		if len(s) > 1 { //当长度大于1的时候，才有必要去特殊罗马数字map中查找
			chars := s[0:2] //首先我们得拿出两个字符去特殊的map中查找
			if v, ok := specialRomanStringMap[chars]; ok { //当存在的时候记录值
				result += v
				s = s[2:]
			} else { //不存在的时候去正常map中查找，并记录
				result += romanStringMap[string(s[0])]
				s = s[1:]
			}
		} else { //当字符串的长度小于等于1的就只能去正常的罗马数字map中查找
			result += romanStringMap[string(s[0])]
			s = s[1:]
		}
	}
	return result
}

func romanToInt1(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]//拿到最后一位

		sign := 1//用于标记是减还是加
		if temp < last {
			//小数在大数的左边，要减去小数
			sign = -1
		}

		res += sign * temp

		last = temp
	}

	return res
}

func main()  {
    fmt.Println(romanToInt("MIIIM"))
    fmt.Println(romanToInt2("MIIIM"))
    fmt.Println(romanToInt1("MIIIM"))
}