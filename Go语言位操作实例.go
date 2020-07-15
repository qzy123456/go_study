package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

// 获取0-n之间的所有偶数
func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return array
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1// Go语言取反方式和C语言不同，Go语言不支持~符号。
}



const cl  = 100

var bl    = 123

type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	return h(name) + 10
}

func main() {
	fmt.Printf("even: %v\n", even(100))
	a, b := swap(100, 200)
	fmt.Printf("swap: %d\t%d\n", a, b)
	fmt.Printf("shifting: %d\n", shifting(100))
	fmt.Printf("nagation: %d\n", nagation(-100))
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1,s2...)
	fmt.Println(s1)
	println(&bl,bl)
	//println(&cl,cl)  //编译不通过，因为常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，
	type MyInt1 int
	type MyInt2 = int
	var i int =9
	var i1 MyInt1 =MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1,i2)


	aa := [3]int{1, 2, 3}
	for k, v := range aa {
		if k == 0 {
			aa[0], aa[1] = 100, 200
			fmt.Print(aa)
		}
		aa[k] = 100 + v
	}
	fmt.Print(aa)
	//[100 200 3][101 102 103]
	ab := []int{1, 2, 3}
	for k, v := range ab {
		if k == 0 {
			ab[0], ab[1] = 100, 200
			fmt.Print(ab)
		}
		ab[k] = 100 + v
	}
	fmt.Print(ab)
	fmt.Println()
   //[100 200 3][101 300 103]
   //这里的关键就是range 后面的那个array或者是slice是复制了上面了的复制值，
   // 所以使用array的话，这个复制的a的v一直都是1 2 3
   // 但是slice的话 就算是复制值那么复制的也是一个 struct{&ptr len cap}所以底层的array可是被改变了，
   // 那么v也是被改变了，所以它的v就成了 1 200 3 所以最后就变成了 101 300 103
   //奇怪的是  只有k == 0 才会这个样子
	var hello handler = func(name string) int {
		return 666
	}
	result := hello.add("还有这种操作")
	fmt.Println(result)
	char := "♥"
	fmt.Println(utf8.RuneCountInString(char))	// 1
	var d uint8 = 2
	fmt.Printf("%08b\n", d)		// 00000010
	fmt.Printf("%08b\n", ^d)		// 11111101

	var data = []byte(`{"status": 200}`)
	var result11 struct {
		Status uint64 `json:"status"`
	}

	 json.NewDecoder(bytes.NewReader(data)).Decode(&result11)

	fmt.Printf("Result: %+v \n", result11)

	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		if err != nil{
			fmt.Println(err)
		}

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
	var byteC byte = 'j'
	fmt.Printf("字符 %c 对应的整型为 %d\n", byteC, byteC)
	defer func() {
		// 2. before : result = 10
		fmt.Printf("before : result = %v\n", result)

		result++

		// 3. after : result = 11
		fmt.Printf("after : result = %v\n", result)
	}()

	result = 10

	// 1. return : result = 10
	fmt.Printf("return : result = %v\n", result)

}


