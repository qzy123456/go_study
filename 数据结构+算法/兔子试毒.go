package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/*
* @Author: polo.huang
* @Date: 2021-08-18 09:50:58
 * @LastEditors: polo.huang
 * @LastEditTime: 2021-08-18 16:50:41
* @Version: 1.0
* @Description: 每日一题
*/
// 1. 实验准备
//1:只追求时间复杂度，同时使用1000只兔纸试毒，这样1天就可以试出毒药，不过就会消耗最多的兔纸
//2:只追求空间复杂度，只用一只兔纸试毒，每天试一瓶，这样只要兔纸死了就找到了毒药，不过这样最坏的结果就是需要1000天才能知道结果
//3:只追求时间也追求空间，把1000分两份试，试完再把500分两份，250分两份依次分到最后
//这样循环迭代，1000->500->250->125->63->32->16->8->4->2->1
//10个箭头，就是10次，也就是用了10天。这样我们只用了10只兔子
//4:利用二进制思路解决问题
//有没有毒，是一个布尔值，也就是可以理解为0为无毒1为有毒
//所以把1-1000瓶药水用二进制编码可以设为 0000000001 - 1111101000
//1111101000刚好是十位数然后用十只兔纸排好顺序喝下对应位数为1编号的药水
//这样只需要1天时间就可以知道十位数所能表达的的二进制数转为就是1023
//也就是说，1天时间内这十只兔纸可以有1023种死法，那根据对应死法的死兔纸的编号可以得出这个二进制数转为十进制数也就是第几瓶药水有毒了
//并且这种算法不仅能算出1000瓶药水有毒，最多可以算出1024瓶药水有毒
//因为，所有兔纸按照这种算法可以喝过1023瓶药水有毒，照这样只要1天后如果没有兔纸死，那就是第1024瓶药水有毒
//————————————————
//版权声明：本文为CSDN博主「半本说明书」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
//原文链接：https://blog.csdn.net/weixin_36825982/article/details/119784215
//    1. 准备1000瓶药水
var potionCount = 1000
var potions = make(map[string]bool) // 索引为字符串代表编号，值为布尔值，true为有毒false为无毒

//    2. 准备一个兔纸池
var rabbits []bool // 索引代表编号，值为布尔值，false为活着true为死了
// 抓兔纸
func grab(count int) {
	rabbits = append(rabbits, make([]bool, count)...) // 按照指定数量抓取新的兔纸放在兔纸池里
}

//    3. 准备一个实验池，可以记录哪只兔纸喝了哪瓶药水，每天结束后根据毒药水编号毒死对应的兔纸
var experiment = make(map[string]int) // 索引为药水编号，值为兔纸的编号
// 兔纸喝药水
func drink(rabbitNum int, potionNum string) {
	experiment[potionNum] = rabbitNum
}
//第一种算法：1kr-1d
func kr1d() {
	// 抓来1000只兔纸
	grab(potionCount)
	// 随机生成毒药水编号
	// fmt.Printf("test: %v\n", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	poisonNum := rand.Intn(potionCount - 1)
	// fmt.Printf("毒药水编码：%v\n", poisonNum)
	// 给药水编号
	for i := 0; i < potionCount; i++ {
		if i == poisonNum {
			potions[strconv.Itoa(i)] = true
		} else {
			potions[strconv.Itoa(i)] = false
		}
	}
	// fmt.Printf("药水：%v\n", potions)
	// fmt.Printf("兔纸：%v\n", rabbits)
	// 开始实验
	// 通过循环代表天数，一次循环代表一天
	fmt.Println("实验开始")
	fmt.Printf("药水数量为：%v\n", len(potions))
	fmt.Println("...")
	flag := false // 为true时结束实验
	for day := 1; true; day++ {
		// fmt.Printf("现在是第%v天\n", day)
		// 给兔纸喝药
		for num, _ := range rabbits {
			potionNum := strconv.Itoa(num) // 根据相同数量的每只兔纸的编号喝对应编号的药水
			drink(num, potionNum)
		}
		// fmt.Printf("全部喝完药水后的实验池：%v\n", experiment)
		// 一天结束了，现在确定实验池中的兔纸谁死了
		for potionNum, rabbitNum := range experiment {
			if potions[potionNum] { // true为有毒false为无毒
				rabbits[rabbitNum] = true // 确认当前药水有毒，毒死兔纸
				fmt.Printf("实验结束\n使用了%v只兔纸 花费了%v天\n编号为%v的兔纸死了 有毒的药水编号为：%v\n",
					len(rabbits), day, rabbitNum, potionNum)
				flag = true
			}
		}
		// 结束实验
		if flag {
			break
		}
		// 结束实验
		if day > potionCount {
			fmt.Printf("实验结束\n超过了1000天找不出毒药中断退出\n")
			break
		}
	}
}
//1r-1~1000d
func r1000d() {
	// 抓来1只兔纸
	grab(1)
	// 随机生成毒药水编号
	// fmt.Printf("test: %v\n", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	poisonNum := rand.Intn(potionCount - 1)
	// poisonNum = potionCount - 1 // 假设第1000瓶药水有毒
	// fmt.Printf("毒药水编码：%v\n", poisonNum)
	// 给药水编号
	for i := 0; i < potionCount; i++ {
		if i == poisonNum {
			potions[strconv.Itoa(i)] = true
		} else {
			potions[strconv.Itoa(i)] = false
		}
	}
	// fmt.Printf("药水：%v\n", potions)
	// fmt.Printf("兔纸：%v\n", rabbits)
	// 开始实验
	// 通过循环代表天数，一次循环代表一天
	fmt.Println("实验开始")
	fmt.Printf("药水数量为：%v\n", len(potions))
	fmt.Println("...")
	flag := false // 为true时结束实验
	for day := 1; true; day++ {
		// fmt.Printf("现在是第%v天\n", day)
		// 给兔纸按照天数喝对应编号的药水
		drink(0, strconv.Itoa(day-1))
		// fmt.Printf("全部喝完药水后的实验池：%v\n", experiment)
		// 一天结束了，现在确定实验池中的兔纸谁死了
		for potionNum, rabbitNum := range experiment {
			// fmt.Printf("test: %v\n", rabbitNum)
			if potions[potionNum] { // true为有毒false为无毒
				rabbits[rabbitNum] = true // 确认当前药水有毒，毒死兔纸
				fmt.Printf("实验结束\n使用了%v只兔纸 花费了%v天\n编号为%v的兔纸死了 有毒的药水编号为：%v\n",
					len(rabbits), day, rabbitNum, potionNum)
				flag = true
			}
		}
		// 清空实验
		experiment = make(map[string]int)
		// 结束实验
		if flag {
			break
		}
		// 结束实验
		if day > potionCount {
			fmt.Printf("实验结束\n超过了1000天找不出毒药中断退出\n")
			break
		}
	}
}
//10r-10d
func r1010d()  {
	// 随机生成毒药水编号
	// fmt.Printf("test: %v\n", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	poisonNum := rand.Intn(potionCount - 1)
	// fmt.Printf("毒药水编码：%v\n", poisonNum)
	// 给药水编号
	for i := 0; i < potionCount; i++ {
		if i == poisonNum {
			potions[strconv.Itoa(i)] = true
		} else {
			potions[strconv.Itoa(i)] = false
		}
	}
	// fmt.Printf("药水：%v\n", potions)
	// 开始实验
	// 通过循环代表天数，一次循环代表一天
	fmt.Println("实验开始")
	fmt.Printf("药水数量为：%v 接下来我们每天平分药水分为A|B两份\n", len(potions))
	fmt.Println("...")
	flag := false // 为true时结束实验
	_potionCount := potionCount
	potionStart := 0
	var halfPotionCount int
	for day := 1; true; day++ {
		// fmt.Printf("现在是第%v天\n", day)
		fmt.Printf("实验第%v天\n", day)
		// 每天把药水平均分为两份每瓶药水只取一滴，然后抓来一只兔纸来喝完其中一份。也就是每天抓一只兔纸
		if _potionCount%2 == 0 {
			halfPotionCount = _potionCount/2 + potionStart
		} else {
			halfPotionCount = _potionCount/2 + 1 + potionStart
		}
		_potionCount = _potionCount / 2
		var potionsA = make(map[string]bool)
		var potionsB = make(map[string]bool)
		// fmt.Printf("_potionCount: %v\n", _potionCount)
		// fmt.Printf("halfPotionCount: %v\n", halfPotionCount)
		for num, potion := range potions {
			_num, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("fail to atoi: %v\n", err)
				_num = 0
			}
			// if day == 2 {
			//     fmt.Printf("test: %v\n", _num)
			// }
			if _num < halfPotionCount {
				potionsA[num] = potion
			} else {
				potionsB[num] = potion
			}
		}
		// fmt.Printf("potionsA: %v\n", potionsA)
		// fmt.Printf("potionsB: %v\n", potionsB)
		fmt.Printf("现在平分药水A: %v瓶 B: %v瓶\n", len(potionsA), len(potionsB))
		// 抓兔纸
		grab(1)
		rabbitNum := day - 1 //每天抓一只兔纸，所以这里以天数为兔纸的编号
		// 给兔纸喝水
		for num, _ := range potionsA {
			drink(rabbitNum, num)
		}

		// fmt.Printf("全部喝完药水后的实验池：%v\n", experiment)
		// 一天结束了，现在确定实验池中的兔纸谁死了
		for potionNum, rabbitNum := range experiment {
			if potions[potionNum] { // true为有毒false为无毒
				rabbits[rabbitNum] = true // 这里只能确定当前这份药水有毒，毒死兔纸
			}
		}
		// 清空实验
		experiment = make(map[string]int)
		if len(potionsA) <= 1 || (!rabbits[rabbitNum] && len(potionsB) <= 1) {
			// 当前这份药水只剩一瓶则结束实验确认毒药水
			flag = true
			var posionNum int
			if rabbits[rabbitNum] {
				// 如果兔纸死了，则这份药水中含有毒药
				posionNum = potionStart
			} else {
				// 反之则在另一份药水
				posionNum = potionStart + 1
			}
			fmt.Printf("实验结束\n使用了%v只兔纸 花费了%v天\n有毒的药水编号为：%v\n", len(rabbits), day, posionNum)
		} else {
			// 否则明天天继续平分这份药水
			if rabbits[rabbitNum] {
				fmt.Printf("兔纸死了所以毒药在A这里\n")
				// 如果兔纸死了，则这份药水中含有毒药
				potions = potionsA
			} else {
				fmt.Printf("兔纸没死所以毒药在B那里\n")
				potionStart = halfPotionCount
				// 反之则在另一份药水
				potions = potionsB
			}
		}
		// fmt.Printf("test: %v\n", len(_potions))
		// 结束实验
		if flag {
			break
		}
		// 结束实验
		if day > 10 {
			fmt.Printf("实验结束\n超过了1000天找不出毒药中断退出\n")
			break
		}
	}
}
//    3. 新的实验池
var _experiment = make(map[int][]string) // 索引为兔纸的编号，值为药水编号的数组
func _drink(rabbitNum int, potionNum string) {
	_experiment[rabbitNum] = append(_experiment[rabbitNum], potionNum)
}

//10r-1d
func r101d()  {
	// 根据药水的数量转二进制的位数抓取对应位数数量的兔纸
	grab(func() (p int) {
		_potionCount := potionCount
		// 获取10进制数的位数
		for ; _potionCount > 0; p++ {
			_potionCount = _potionCount >> 1
			// fmt.Printf("%v\n", potionCount)
			// fmt.Printf("%v\n", p)
		}
		// fmt.Printf("%v\n", p)
		return
	}())
	// 随机生成毒药水编号
	// fmt.Printf("test: %v\n", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	poisonNum := rand.Intn(potionCount)
	// fmt.Printf("毒药水编码：%v\n", poisonNum)
	// 给药水编号
	for i := 1; i <= potionCount; i++ {
		if i == poisonNum {
			potions[strconv.Itoa(i)] = true
		} else {
			potions[strconv.Itoa(i)] = false
		}
	}
	// fmt.Printf("药水：%v\n", potions)
	// fmt.Printf("兔纸：%v\n", rabbits)
	// 开始实验
	// 通过循环代表天数，一次循环代表一天
	fmt.Println("实验开始")
	fmt.Printf("药水数量为：%v\n", len(potions))
	fmt.Println("...")
	flag := false // 为true时结束实验
	for day := 1; true; day++ {
		// fmt.Printf("现在是第%v天\n", day)
		// 给兔纸喝药
		fmt.Println("开始喂兔纸喝药水\n...")
		for rabbitNum, _ := range rabbits {
			for potionNum, _ := range potions {
				// 利用当前编号的兔纸对应的二进制位数判断当前编号的药水在二进制中指定位数是否为1，是则让这只兔纸喝这瓶水
				_potionNum, err := strconv.Atoi(potionNum)
				if err != nil {
					panic(err)
				}
				// if _potionNum < 10 {
				//     fmt.Printf("potionNum: %v rabbitNum: %v res: %v\n", (_potionNum + 1), (rabbitNum + 1), ((_potionNum+1)>>(rabbitNum))&1)
				// }
				if (_potionNum >> rabbitNum & 1) == 1 {
					// 当前编号的药水在二进制中指定位数为1
					_drink(rabbitNum, potionNum)
				}
			}
			fmt.Printf("第%v这只兔纸喝了%v瓶药的水\n", rabbitNum, len(_experiment[rabbitNum]))
		}
		fmt.Println("所有兔纸都喝完了\n...")
		// fmt.Printf("experiment: %v\n", _experiment)
		// 一天结束了，现在确定实验池中的兔纸谁死了
		fmt.Println("毒性发作中\n...")
		var posionNum int // 有毒的药水编号
		for rabbitNum, potionNums := range _experiment {
			for _, potionNum := range potionNums {
				if potions[potionNum] { // true为有毒false为无毒
					rabbits[rabbitNum] = true // 这里只能确定当前这份药水有毒，毒死兔纸
					// 检查兔纸的死亡情况
					fmt.Printf("第%v只兔纸死了\n", rabbitNum)
					// fmt.Printf("rabbitNum: %v\n", rabbitNum)
					posionNum += int(math.Pow(2, float64(rabbitNum)))
					break
				}
			}
		}
		fmt.Println("从右到左打印兔纸的死亡情况：(1为死亡 0为活着) 也就是有毒药水的二进制表示")
		for i := len(rabbits); i > 0; i-- {
			if rabbits[i-1] {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
		// fmt.Printf("rabbits: %v\n", rabbits)
		// fmt.Printf("posionNum: %v\n", posionNum)
		// 判断有毒药水是否正确
		if posionNum > 0 && posionNum <= potionCount {
			flag = true
			fmt.Printf("实验结束\n使用了%v只兔纸 花费了%v天\n有毒的药水编号为：%v\n", len(rabbits), day, posionNum)
		}

		// 结束实验
		if flag {
			break
		}
		// 结束实验
		if day > potionCount {
			fmt.Printf("实验结束\n超过了1000天找不出毒药中断退出\n")
			break
		}
	}
}
func main() {
	kr1d()
	r1000d()
	r1010d()
	r101d()
}
