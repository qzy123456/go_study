package main

import "fmt"
func main() {
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])  //999
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1]) // 0 999
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i]) //4 5 2 3 4
	}
	wrongCopyMatrix()
	rightCopyMatrix()
	//例如
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

}
//坑位一：切片 dst 需要先初始化长度
//不是你定义好类型，就能将 src 完全 copy 到 dst 的，你需要初始化长度。
//
//如果 dst 长度小于 src 的长度，则 copy 部分；
//如果大于，则全部拷贝过来，只是没占满 dst 的坑位而已；
//相等时刚好不多不少 copy 过来。
//坑位二：源切片中元素类型为引用类型时，拷贝的是引用
//由于只 copy 切片中的元素，所以如果切片元素的类型是引用类型，那么 copy 的也将是个引用。
//
//如下面例子，matA 和 matB 地址不一样，但 matA[0] 和 matB[0] 的地址是一样的。
func wrongCopyMatrix() {
	matA := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}
	matB := make([][]int, len(matA))
	copy(matB, matA)
	fmt.Printf("%p, %p\n", matA, matA[0]) // 0xc0000c0000, 0xc0000c2000
	fmt.Printf("%p, %p\n", matB, matB[0]) // 0xc0000c0050, 0xc0000c2000
}
//如果想 copy 多维切片中的每一个切片类型的元素，那么你需要将每个切片元素进行 初始化 并 拷贝。注意是两步：先 初始化，再 拷贝。
//
//正确的拷贝一个多维数组：
func rightCopyMatrix() {
	matA := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}
	matB := make([][]int, len(matA))
	for i := range matA {
		matB[i] = make([]int, len(matA[i])) // 注意初始化长度
		copy(matB[i], matA[i])
	}
	fmt.Printf("%p, %p\n", matA, matA[0]) // 0xc00005c050, 0xc000018560
	fmt.Printf("%p, %p\n", matB, matB[0]) // 0xc00005c0a0, 0xc0000185c0
}