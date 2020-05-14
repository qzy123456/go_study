package main

import (
	"fmt"
	"strings"
)

func main()  {
	b := strings.EqualFold("Home", "home")  //验证字符串是否相等  不区分大小写
	fmt.Println(b)
	b1 := strings.HasPrefix("Home", "h") //是不是某个字符串开头  区分大小写
	fmt.Println(b1)
	b2 := strings.HasSuffix("Home", "me")  //是不是某个字符串结尾  区分大小写
	fmt.Println(b2)
	b3 := strings.Contains("HOME", "ho")  //判断一个字符串是否在某个字符串中～区分大小写
	fmt.Println(b3)
	position := strings.Index("lorem lorem", "lo") //判断一个字符串在某个字符串首次出现的位置  不存在返回 -1
	fmt.Println(position)
	num := strings.IndexFunc("me", func(r rune)bool {  //判断一个字符串在某个字符串首次出现的位置（方法）  不存在返回 -1
		return r == rune('m')
	})
	fmt.Println(num)
	num1 := strings.LastIndex("mm", "m") //返回字符串，最后一次出现的位置
	fmt.Println(num1)

	str := strings.Title("go home") //首字母大写
	fmt.Println(str)

	str1 := strings.ToLower("GO HOME") //转成小写
	fmt.Println(str1)

	str2 := strings.ToUpper("go home") //转成大写
	fmt.Println(str2)

	str3 := strings.Repeat("m", 3) //重复  替换
	fmt.Println(str3)
	/*
    参数
    [string] 被处理字符
    [string] 匹配字符
    [string] 替换字符
    [int] 替换个数
	*/
	str4 := strings.Replace("co co co co", "co", "jc", -1)//替换
	fmt.Println(str4)

	str5 := strings.Trim(" - title - ", "-")//去除连端的 某个字符
	fmt.Println(str5)

	str6 := strings.TrimSpace("  titie  ") //去除两端的空格
	fmt.Println(str6)

	str7 := strings.Fields("coco jeck ")//按照空格  切割字符串
	fmt.Println(str7)

	// 使用逗号分割
	str8 := strings.FieldsFunc("coco,jeck,andy", func(r rune) bool {  // 使用逗号分割
		return r == rune(',')
	})
	fmt.Println(str8)

	str9 := strings.Split("product/id/place", "/") //使用指定字符作为分割符   切割字符串
	fmt.Println(str9)

	str10 := strings.SplitN("product/id/place", "/", 4) //指定切分数量的Split  也就是切割成几分（大于那个值  只取最大值）
	fmt.Println(str10)

	str11 := strings.Join([]string{"coco", "jeck"}, ",")//合并字符串
	fmt.Println(str11)

	str12 := strings.Count("hello word","o") //字符串出现的次数
	fmt.Println(str12)


}
