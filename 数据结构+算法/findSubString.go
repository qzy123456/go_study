package main

import "fmt"

func findSubstring(s string, words []string) []int {
	// Get word length
	wordLength := len(words[0])
	wordNum := len(words)
	stringLength := len(s)
	var ans []int
	// No words
	if wordNum == 0 {
		return ans
	}
	if stringLength < wordLength {
		return ans
	}
	// The number of words should appear
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	//fmt.Printf("%#v",wordMap) //{"bar":1, "foo":1}
	// Building sliding windows
	for left := 0; left <= stringLength-wordLength*wordNum; left++ {
		// Create a map Record the occurrence of words in the sliding window
		tempMap := make(map[string]int)
		flag := true // Record whether the substring is legal
		for right := left; right < left+wordLength*wordNum; right += wordLength {
			// Intercept a word
			temp := s[right : right+wordLength]
			if wordMap[temp] == 0 {
				// 没有这个词
				flag = false
				break
			}
			// There is this word , It was recorded that map in
			tempMap[temp]++
			if wordMap[temp] < tempMap[temp] {
				// Too many words （ Too few may appear later ）
				flag = false
				break
			}
		}
		if flag {
			// Pass the test
			ans = append(ans, left)
		}
	}
	return ans
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	//s1 := "wordgoodgoodgoodbestword"
	//words1 := []string{"word", "good", "best", "word"}

	fmt.Println(findSubstring(s, words))
	//fmt.Println(findSubstring(s1, words1))
}
