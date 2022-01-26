package main

import "fmt"

//由于字母共有 26 个，所以我们可以声明一个 26 个长度的数组（该种方法在本类题型很常用）
// 因为字符串中字母可能是重复的，所以我们可以先进行第一次遍历，在数组中记录每个字母的最后一次出现的所在索引。
// 然后再通过一次循环，比较各个字母第一次出现的索引是否为最后一次的索引。如果是，我们就找到了我们的目标，
// 如果不是我们将其设为 -1（标示该元素非目标元素）如果第二次遍历最终没有找到目标，直接返回 -1即可。
func firstUniqueChar(strs string) int{
	var arr [26]int
	for key, value := range strs {
		arr[value - 'a'] = key
	}

	for key, value := range strs {
		if arr[value - 'a'] == key{
			return key
		}
	}


	return -1
}

func main()  {
     str := "loveleetcode"
     fmt.Println(firstUniqueChar(str))
}