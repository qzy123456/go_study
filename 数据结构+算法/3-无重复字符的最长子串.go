package main

import (
	"fmt"
	"math"
	"strings"
)

//abcabcbb
//判断子串是否合法
//内嵌一个循环判断是否有重复子串
//如果有个重复子串，则start和end的位置增加
//如果不重复则，继续增加子串的数量
func lengthOfLongestSubstring(s string) int {
	if s == ""{
		return 0
	}
	var strLen = len(s)
	var start,end=0,1
	var s2 string
	for end <= strLen{
		var s1 string = string(s[start:end])
		//计算重复
		var repeatCount int = 1
		var s1Len int = len(s1)
		for i:=0;i < s1Len;i++{
			var temp string = string(s1[i])
			repeatCount = strings.Count(s1,temp)
			//说明有重复的了
			if repeatCount >1{
				start++
				end =start+len(s1)
				break
			}
		}
		if repeatCount ==1{
			//增加
			s2 = s1 //保存之前的值
			end++
		}
	}
	return len(s2)
}
//n*n*n  找出该串的所有子串，不重复，且最大即可
func lengthOfLongestSubstring2(s string)int{
	var counts int = 0
	for i:=0;i<len(s);i++{
		//找出他所有的子串
		for j:=i+1;j<=len(s);j++{
			if allUnique(s,i,j){
				temp := math.Max(float64(counts),float64(j-i))
				counts = int(temp)
			}
		}
	}
	return counts
}

//找出子串不重复的
func allUnique(s string,start,end int) bool {
	//会截取start -> end-1位置的下标
	s = s[start:end]
	for i := 0; i < len(s); i++ {
		//返回该子串在主串中的重复数量
		repeatCount := strings.Count(s, string(s[i]))
		if repeatCount != 1 {
			return false
		}
	}
	return true
}

//利用滑动窗口和容器
//使用一个HashSet来实现滑动窗口，用来检查重复字符。 维护开始和结束两个索引，默认都是从0开始，
//然后随着循环【向右移动结束索引】，遇到不是重复字符则放入窗里，遇到重复字符则【向右侧移动开始索引】，最终得到结果
func lengthOfLongestSubstring3(s string)int {
	//借助一个容器，来判断，子串中是否有重复
	m := make(map[byte]int)
	sLen := len(s)
	start := 0
	end   := 0
	repeatCount  := 0
	//start 和 end 双条件判断，只有end一个也可以，可能这样更严谨一些吧
	for start < sLen && end < sLen {
		temp := s[end]
		if _, ok := m[temp]; !ok {
			//不存在说明该key是唯一的
			m[s[end]] = end
			end++ //移动滑动窗口
			repeatCount = max(repeatCount, end-start)
		} else {
			//说明了有重复的，滑动窗口移动，则start+1,
			//这个时候s[start]，s[end]就是第一次出现该字符的位置
			delete(m, s[end])
			start++
		}
		//fmt.Println(m)
	}
	return repeatCount
}
//优化滑动窗口
// abcabcbb
// abc
// bca
// cab
// abc
// cb
// b
func lengthOfLongestSubstring4(s string)int {
	var ans ,index =  0 ,0
	m := make(map[byte]int) //存放字符出现的位置
	for j:= 0; j < len(s); j++ {
	//abcabcbb
	//map[]map[97:1]map[97:1 98:2]map[97:1 98:2 99:3]map[97:4 98:2 99:3]map[97:4 98:5 99:3]map[97:4 98:5 99:6]map[97:7 98:5 99:6]
		if tempIndex, ok := m[s[j]]; ok {
			//发现重复的，则重新选择一个index，这个index停留在出现重复的前一位置
			index = max(tempIndex, index)//这里的index则为下标
		}
		//每次计算j-index+1的记录，j为字符串的下标，index为下标，重复，index则从重复位置前一位开始
		ans = max(ans, j-index+1)
		m[s[j]] = j + 1
	}
	return ans
}

// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring5(s string) int {
	right, left, res := 0, 0, 0
	maps := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := maps[s[left]]; ok && idx >= right {
			right = idx + 1
			fmt.Println(maps[s[left]],string(s[left]),right)
		}
		maps[s[left]] = left
		left++
		fmt.Println("left - right",left -right)
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "abcabcaa"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstring2(s))
	fmt.Println(lengthOfLongestSubstring3(s))
	fmt.Println("4===》",lengthOfLongestSubstring4(s))
	fmt.Println(lengthOfLongestSubstring5(s))
}