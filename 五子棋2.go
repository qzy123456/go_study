package main


import (
	"sync"
   "fmt"
 )

var qipan [][]int
var result = make([]bool, 0, 0)

func showQipan(){
	for i := 0; i<len(qipan); i++{
		for j := 0; j< len(qipan[0]);j++{
			fmt.Printf("%d ", qipan[i][j])
		}
		fmt.Println()
	}
}

func luozi(x, y, color int) {

	var wg sync.WaitGroup
	qipan[x][y] = color

	showQipan() //调试打印

	for i := range direct{
		wg.Add(1)
		go func(d []int){
			v := worker(x,y,color, 0, d,judgeValid)
			result = append(result, v)
			wg.Done()
		}(direct[i])
	}

	wg.Wait()
}

func main(){
	qipan = [][]int{
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0,0},
		{0,0,0,0,1,0,0,0},
		{0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,1,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
	}

	luozi(1, 2, 1)
	fmt.Println(result)
}


var direct = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var num int = 5

func worker(x, y, color, level int,direct []int,f func(x, y, color int) bool) bool{

	if level == num {
		return true
	}

	if !f(x, y, color){
		return false
	}

	return  worker(x+direct[0], y+direct[1], color, level + 1,direct, f)
}


func judgeValid(x, y, color int) bool {
	if x < 0 || x >= len(qipan) || y < 0 || y >= len(qipan[0]) {
		return false
	}

	if qipan[x][y] == color {
		return true
	}

	return false
}

