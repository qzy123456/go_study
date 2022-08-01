package main

import "fmt"

//438. 找到字符串中所有字母异位词
//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//示例 1:
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 示例 2:
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

//滑动窗口
func findAnagrams(s, p string)  []int {
	slen,pLen := len(s),len(p)
	if slen < pLen{
		return []int{}
	}
	var sCount, pCount [26]int
	for i,v := range p{
		pCount[v-'a']++
		sCount[s[i]-'a']++
	}
	ans := make([]int,0)
	if pCount == sCount{
		ans = append(ans,0)
	}

	for i,v :=range s[:slen-pLen]{
		sCount[v-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount{
			ans = append(ans,i+1)
		}
	}
	return ans
}
//优化的滑动窗口
func findAnagrams1(s string, p string) []int {
	var freq [256]int
	var result []int
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}
	left, right, count := 0, 0, len(p)

	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}

	}
	return result
}

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s,p))
}

