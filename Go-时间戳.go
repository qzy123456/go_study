package main
import "time"
import "fmt"
func main() {
	var tt int64 = time.Now().Unix()
	var s string = time.Unix(tt, 0).Format("2006-01-02 15:04:05")
	println("我是时间戳转年月日",s) //打印出 我是时间戳转年月日 2019-12-04 18:33:29

	toBeCharge := s  //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写，不然对应不上，会转成负数
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")  //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()  //转化为时间戳 类型是int64
	fmt.Println(theTime)  //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println("我是年月日转时间戳",sr)    //打印输出时间戳 1420041600

}
