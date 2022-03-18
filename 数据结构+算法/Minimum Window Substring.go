package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}

func minWindow2(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	goalMap := make(map[string]int)
	for _, v := range t {
		goalMap[string(v)]++
	}

	res, minLen := "", math.MaxInt32
	l, cnt, golaSize := 0, 0, len(t)
	currMap := make(map[string]int)
	for r := 0; r < len(s); r++ {
		key := string(s[r])
		if goalMap[key] <= 0 {
			continue
		}

		if currMap[key] < goalMap[key] {
			cnt++
		}
		currMap[key]++
		if cnt == golaSize {
			k := string(s[l])
			for goalMap[k] <= 0 || currMap[k] > goalMap[k] {
				if goalMap[k] > 0 && currMap[k] > goalMap[k] {
					currMap[k]--
				}
				l++
				k = string(s[l])
			}
			if r-l+1 < minLen {
				minLen = r - l + 1
				res = s[l : r+1]
			}
		}

	}
	return res
}

func minWindow3(s string, t string) string {
	var res string
	cnt := math.MaxInt32
	hashMap := make(map[byte]int)
	l := 0
	r := 0
	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}

	for r < len(s) {
		hashMap[s[r]]--
		for check(hashMap) {
			if r-l+1 < cnt {
				cnt = r - l + 1
				res = s[l : r+1]
			}
			hashMap[s[l]]++
			l++
		}
		r++
	}
	return res
}

func check(hashMap map[byte]int) bool {
	for _, v := range hashMap {
		if v > 0 {
			return false
		}
	}
	return true
}


func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	fmt.Println(minWindow3(S,T))
}