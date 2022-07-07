package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordLength := len(words[0]) //单词长度，单词长度一致
	wordNum := len(words) //单词数量
	stringLength := len(s) //总词的长度
	var ans []int //返回值

	if wordNum == 0 {
		return ans
	}
	if stringLength < wordLength {
		return ans
	}
	//map存储单词出现的数量
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	//fmt.Printf("%#v",wordMap) //{"bar":1, "foo":1}
	//滑动窗口  总长 - 单词数量 * 单词长度
	for left := 0; left <= stringLength-wordLength*wordNum; left++ {
		// 创建一个map记录单词在滑动窗口中出现次数
		tempMap := make(map[string]int)
		flag := true //有没有找到
		for right := left; right < left + wordLength*wordNum; right += wordLength {
			//滑动取出一个词
			temp := s[right : right+wordLength]
			if wordMap[temp] == 0 {
				// 没有这个词
				flag = false
				break
			}
			// 记录找到了
			tempMap[temp]++
			//出现次数太多
			if wordMap[temp] < tempMap[temp] {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, left)
		}
	}
	return ans
}
//官方答案
func findSubstring2(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
	fmt.Println(findSubstring2(s, words))
}
