package main

import (
	"bytes"
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}



func diff(s1, s2 string) int {
	cnts := [26]int{}
	for _, c := range s1 { cnts[c-'a']++ }
	for _, c := range s2 { cnts[c-'a']-- }
	for _, cnt := range cnts {
		if cnt != 0 {
			return cnt
		}
	}
	return 0
}

func groupAnagrams2(strs []string) (res [][]string) {
	//计数的方式直接对字符串先排序
	sort.Slice(strs, func(i, j int) bool {
		return diff(strs[i], strs[j]) < 0
	})

	for i, str := range strs {
		if i == 0 || diff(str, strs[i-1]) != 0 {
			res = append(res, nil)
		}
		res[len(res)-1] = append(res[len(res)-1], str)
	}
	return
}


func getKey(str string) string {
	// 记录 26 个英文字母出现的次数
	cnt := [26]int{}
	for _, c := range str {
		cnt[int(c-'a')]++
	}
	// 各个字母出现次数相同即为一组，拥有相同的 key
	var key bytes.Buffer
	for _, n := range cnt {
		key.WriteByte(byte(n))
	}
	return key.String()
}

func groupAnagrams3(strs []string) [][]string {
	var result [][]string
	kv := map[string][]string{}

	for _, s := range strs {
		key := getKey(s)
		kv[key] = append(kv[key], s)
	}

	for _, v := range kv {
		result = append(result, v)
	}
	return result
}

func groupAnagrams4(strs []string) [][]string {
	//定义map存储相同字符
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		//把排序过后，字母一样的字符串存储到一起，类似于
		//map[abt:[bat] aet:[eat tea ate] ant:[tan nat]]
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	//拼接返回值
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams5(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	strs :=  []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
	fmt.Println(groupAnagrams3(strs))
	fmt.Println(groupAnagrams4(strs))
	fmt.Println(groupAnagrams5(strs))
}