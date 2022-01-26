package main

import "fmt"

func getPrefix(arr []string) string {
	if len(arr) <= 1 {
		return ""
	}
	firstStr := arr[0]
	l := len(arr)
	for i := range firstStr {
		for j := 1; j < l; j++ {
			if arr[j][i] != firstStr[i] {
				return firstStr[:i]
			}
		}
	}
	return ""
}

func main() {
	arr :=[]string{"flower","flow","fli","ff"}
	fmt.Println(getPrefix(arr))
}
