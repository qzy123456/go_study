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
///给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//注意：
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
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
			if r+1-l < cnt {
				cnt = r + 1 - l
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

func minWindow4(s string, t string) string {
	wind := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	left, right, match, start, end, min := 0, 0, 0, 0, 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			wind[c]++
			if wind[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			if need[c] != 0 {
				//可能存在s-->aaaa  t-->a,
				//这里只有当s的最后一个a也被移出窗口的时候
				//匹配数才少了1
				if wind[c] == need[c] {
					match--
				}
				wind[c]--
			}
		}
	}
	if min == math.MaxInt32 {
		return ""
	}
	return s[start:end]
}

func minWindow5(s string, t string) string {
	mps,mpt:=make(map[byte]int),make(map[byte]int)
	for i:=range t{
		mpt[t[i]]++
	}
	ansL,ll,l,r:=-1,math.MaxInt64,0,0

	check:=func()bool{
		for k,v:=range mpt{
			if mps[k]<v{
				return false
			}
		}
		return true
	}

	for ;r<len(s);r++{
		if r<len(s) && mpt[s[r]]>0{
			mps[s[r]]++
		}
		for check()&& l<=r{
			if r-l+1<ll{
				ll=r-l+1
				ansL=l
			}
			if mpt[s[l]]>0{
				mps[s[l]]--
			}
			l++
		}
	}
	if ll==math.MaxInt64{
		return ""
	}
	return s[ansL:ansL+ll]
}

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	fmt.Println(minWindow3(S,T))
}