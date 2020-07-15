package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	//邮箱
	reQQEmail = `(\d+)@qq.com`
	reEmail   = `\w+@\w+\.\w+(\.\w+)?`

	//超链接
	//<a href="http://news.baidu.com/ns?cl=2&rn=20&tn=news&word=%C1%F4%CF%C2%D3%CA%CF%E4%20%B5%BA%B9%FA"
	reLinkBad = `<a[\s\S]*?href="(https?://[\s\S]+?)"`
	reLink    = `href="(https?://[\s\S]+?)"`

	//手机号
	//13x xxxx xxxx
	rePhone = `1[345789]\d\s?\d{4}\s?\d{4}`

	//身份证号
	//123456 1990 0817 123X
	reIdcard = `[123456]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dX]`

	//图片链接
	//"http://img2.imgtn.bdimg.com/it/u=2403021088,4222830812&fm=26&gp=0.jpg"
	reImg = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(svg)|(swf)|(ico))))"`
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Print(why, err)
	}
}
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr = string(pageBytes)
	return pageStr
}
func SpiderEmail() {
	pageStr := GetPageStr("http://tieba.baidu.com/p/2544042204")
	pageStr += "ximendong@21centry.com.cn"
	//fmt.Println(pageStr)

	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		//fmt.Printf("email=%s,qq=%s\n",result[0],result[1])
		fmt.Println(result)
	}
}
func SpiderLink() {
	pageStr := GetPageStr("http://www.baidu.com/s?wd=%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1%20%E5%B2%9B%E5%9B%BD")
	//fmt.Println(pageStr)

	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果:\n", len(results))
	for _, result := range results {
		//fmt.Printf("email=%s,qq=%s\n",result[0],result[1])
		fmt.Println(result[1])
		fmt.Println()
	}
}

func SpiderMobilePhone() {
	pageStr := GetPageStr("http://www.zhaohaowang.com/aspx/zhw/index.html?CityId=1#BJ=05698")
	//fmt.Println(pageStr)

	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果:\n", len(results))
	for _, result := range results {
		fmt.Println(result)
		fmt.Println()
	}

}
func SpiderIdcard() {
	pageStr := GetPageStr("http://sfz.ckd.cc/")
	//fmt.Println(pageStr)

	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果:\n", len(results))
	for _, result := range results {
		fmt.Println(result[0])
		fmt.Println()
	}

}
func SpiderImg() {
	pageStr := GetPageStr("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
	//fmt.Println(pageStr)

	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果:\n", len(results))
	for _, result := range results {
		fmt.Println(result[1])
		fmt.Println()
	}

}
func main() {
	SpiderImg()
	SpiderIdcard()
	SpiderMobilePhone()
	SpiderEmail()
	SpiderLink()
}
