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

func main()  {
    fmt.Println(romanToInt("MIIIM"))
    fmt.Println(romanToInt("LVIII"))
}