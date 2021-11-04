
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 匹配的结构
type Match_player struct {
	UID int
	Lev int
}

// 匹配的chan
var Match_Chan chan *Match_player
var Imax int = 0

// 初始化
func init() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	re := InitDSQ(data)
	fmt.Println("初始化：",re)

	Match_Chan = make(chan *Match_player, 100)
	return
}
func RandInterval_LollipopGo(b1, b2 int32) int32 {
	if b1 == b2 {
		return b1
	}

	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	return int32(rand.Int63n(max-min+1) + min)
}

func InitDSQ(data1 []int) [4][4]int {

	data := data1
	erdata := [4][4]int{}
	j, k := 0, 0

	// 循环获取
	for i := 0; i < 16; i++ {
		// 删除第i个元素
		//获取随机数
		icount := RandInterval_LollipopGo(0, int32(len(data))-1)
		fmt.Println("随机数：", icount)
		//datatmp := data[icount]

		if len(data) == 1 {
			erdata[3][3] = data[0]
		} else {
			//------------------------------------------------------------------
			//如果随机小于总长度
			if int(icount) < len(data) {
				erdata[j][k] = data[icount]
				k++
				if k%4 == 0 {
					j++
					k = 0
				}

				data = append(data[:icount], data[icount+1:]...)
				fmt.Println(data)
			}
			//------------------------------------------------------------------
		}
		fmt.Println("生成的数据", erdata)
	}

	return erdata
}

// 主函数
func main() {

	// 第一个数据：
	idata := &Match_player{
		UID: 1,
		Lev: 6,
	}
	Putdata(idata)

	// 第二个数据：
	idata1 := &Match_player{
		UID: 2,
		Lev: 20,
	}
	Putdata(idata1)

	// 第三个数据：
	idata2 := &Match_player{
		UID: 3,
		Lev: 90,
	}
	Putdata(idata2)

	// 第四个数据：
	idata3 := &Match_player{
		UID: 3,
		Lev: 900,
	}
	Putdata(idata3)
	Putdata(idata3)
	Putdata(idata3)

	// defer close(Match_Chan)
	Imax = len(Match_Chan)
	// 取数据
	 DoingMatch()
	 go Sort_timer()

	strport := "8892" //  GM 系统操作 -- 修改金币等操作
	http.ListenAndServe(":"+strport, nil)

	return
}

// 压入
func Putdata(data *Match_player) {
	Match_Chan <- data
	return
}

// 获取
func DoingMatch() {

	Data := make(map[int]*Match_player)
	// 全部数据都拿出来
	// data := make(chan map[string]*Match_player, 100)
	// data <- Match_Chan
	for i := 0; i < Imax; i++ {
		if data, ok := <-Match_Chan; ok {
			fmt.Print(data, "\t")
			Data[i+1] = data
		} else {
			fmt.Print("woring", "\t")
			break
		}
	}
	// 打印数据保存
	fmt.Println(Data)
	return
}

func Sort_timer() {
	// 控制排队的速度
	timer := time.NewTimer(time.Millisecond * 400)
	for {
		select {
		case <-timer.C:
			{
				// 获取channel数据的函数。
				DoingMatch()
			}
		}
	}
}

