package main

//链表实现
import (
	"fmt"
	"os"
)

//定义错误常量
const (
	ERROR = -1000000001
)

//定义元素类型
type Elements int64

//定义节点
type LinkNode struct {
	Data Elements   //数据域
	Next *LinkNode //指针域，指向下一个节点
}

//末尾添加数据
//@param head *LinkNode 头节点
//@param data Elements int型数字
func Add(head *LinkNode, data Elements) {
	point := head //临时指针（头指针）
	//找到尾指针
	for point.Next != nil {
		point = point.Next //移位
	}
	//定义新节点
	var node LinkNode  //新节点
	point.Next = &node //赋值
	node.Data = data
}

//删除指定index位置的数据
func Delete(head *LinkNode, index int) Elements {
	//判断index合法性
	//如果下标小于0，或者下标大于链表的长度
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check delete index")
		return ERROR
	} else {
		//临时节点，等于头节点～
		point := head
		//找到这个节点的上一个节点
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		//当前节点，也就是要删除的上个节点。（当前节点的下一个，变成要删除的节点的下一个）
		point.Next = point.Next.Next //赋值
		//数据同理
		data := point.Next.Data
		return data
	}
}

//插入指定index位置的数据
func Inserts(head *LinkNode, index int, data Elements) {
	//检验index合法性
	//判断下标如果小于0 或者下标大于链表的最大长度
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check  insert index")
	} else {
		//临时节点，也就是头节点
		point := head
		//找到要插入下标的上个节点
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		var node LinkNode //新节点，赋值
		//当前节点（也就是要插入节点的上个节点的下一个，变成要插入数据的下一个，上个节点的下一个也就变成了自己）
		node.Data = data
		node.Next = point.Next
		//上个节点的下一个也就变成了自己 (PS：⚠️这里顺序不能变，不然会死循环)
		point.Next = &node
	}
}

//获取长度
func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

//搜索 头结点 data元素
func Search(head *LinkNode, data Elements) {
	//定义头节点，也就是临时节点
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Next
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}

//获取指定节点的值
func GetData(head *LinkNode, index int) Elements {
	point := head
	//判断下标的正确性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check getData index")
		return ERROR
	} else {
		for i := 0; i < index; i++ {
			point = point.Next
		}
		return point.Data
	}
}

//遍历 头结点
func Traverse(head *LinkNode) {
	point := head.Next
	for point.Next != nil {
		fmt.Println("遍历",point.Data)
		point = point.Next
	}
}

//主函数测试
func main() {
	//定义头节点  也就是空节点
	var head *LinkNode = &LinkNode{Data: 0, Next: nil}
	var nodeArray []Elements
	//生成10个节点数据
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Elements(i+1+i*100))
		Add(head, nodeArray[i])
	}
    //删除节点
	Delete(head, 3)  //删除第3个节点
	//查找节点
	Search(head, 2032) //查询数据
	//插入节点
	Inserts(head, 2, 10010) //第n个节点 插入数据（长度不对  就报错）
	//遍历节点
	Traverse(head)                     //遍历数据
	fmt.Println("data is", GetData(head, 6)) //第n个节点
	fmt.Println("length:", GetLength(head))  //长度
	//退出程序
	os.Exit(0)
	//2032 not exist at
	//please check  insert index
	//遍历 1
	//遍历 102
	//遍历 304
	//遍历 405
	//遍历 506
	//遍历 607
	//遍历 708
	//遍历 809
	//data is 607
	//length: 9
}