package main
import "fmt"

type Element interface{}

var NoData Element = 0
//定义链表的struct
type Node struct {
	Data Element   //链表存储的数据
	Next *Node    //下个节点
	Pre  *Node    //上个节点
}
//头节点
type List struct {
	Head   *Node    //头节点
	Length int     //定义双向链表的长度
}

// 创建双向链表（也就是创建一个空的头节点）
func Create() *List {
	//创建一个空链表，也就是头
	head := new(Node)
	return &List{Head: head, Length: 0} //创建一个头节点
}

// 获取链表长度
func (l *List) Len() int {
	return l.Length
}

// 判断是否为空链表
func (l *List) IsEmpty() bool {
	if l.Head.Next == nil && l.Head.Pre == nil {
		return true
	}
	return false
}

// 向链表末尾追加数据
func (l *List) Append(e Element) {
	//把这个新的要插入的数据包装成新的节点～
	node := &Node{Data: e, Next: nil, Pre: nil}
     //头节点
	p := l.Head
	//如果是个空链表，那么直接插入
	if l.IsEmpty() {
		//头节点的下个节点就是新节点
		p.Next = node
		//新节点的上个节点就是头节点
		node.Pre = p
         //链表的长度加1
		l.Length++
		return
	}
    //如果不是空链表，那么就一直找，找到最后一个
	for p.Next != nil {
		p = p.Next
	}
    //当前节点（也就是最后一个节点的下一个，就是最新节点）
	p.Next = node
	//最新节点的上一个就是最后的那个节点
	node.Pre = p
	//链表的长度加1
	l.Length++
	return
}

// 在头部（第一个位置）追加数据
func (l *List) PreAppend(e Element) {
	//头部节点
	p := l.Head
	//把当前数据  包装成一个新的节点
	node := &Node{Data: e, Next: nil, Pre: nil}
    //空链表的话，那么直接插入
	if l.IsEmpty() {
		p.Next = node
		node.Pre = p

		l.Length++

		return
	}

	// 插入节点的 NEXT 指向头节点的 Next（也就是当前节点的下一个，变成了头节点的下一个）
	node.Next = p.Next
	// 头节点的 Next 的 Pre 指向 新插入的节点 （头节点的上一个变成了当前要插入的节点）
	p.Next.Pre = node
	// 头节点的 Next 指向 新插入的节点
	p.Next = node
	// 新插入节点的 Pre 指向头节点
	node.Pre = p

	l.Length++

	return
}

// 在指定位置插入数据
func (l *List) Insert(index int, e Element) {
	//判断链表是否为空  以及index的有效性
	if l.IsEmpty() || index < 0 {
		l.PreAppend(e)

		return
	}
   //如果index大于链表的最大长度，那么就直接插入到链表的尾部
	if index > l.Len() {
		l.Append(e)

		return
	}
    //链表的头部
	p := l.Head
	//把要插入的数据拼装成一个节点的格式
	node := &Node{Data: e, Next: nil, Pre: nil}
    //找到要插入的位置的上一个节点
	for i := 0; i < index-1; i++ {
		p = p.Next
	}

	// 新插入节点的 Next 节点指向 p[index-1]的Next 节点
	node.Next = p.Next
	// p[index - 1]的 Next.Pre 节点 指向 node 节点
	p.Next.Pre = node

	// p[index -1] 的Next 节点 指向 新插入的节点
	p.Next = node
	// ❤新插入的节点的Pre 指向 p[index-1]
	node.Pre = p
    //链表的长度加1
	l.Length++

	return
}

// 删除指定位置的数据, 并返回该数据
func (l *List) Delete(index int) Element {
	if l.IsEmpty() {
		fmt.Println("list is empty. delete error")
		return NoData
	}

	if index < 0 || index > l.Len() {
		fmt.Println("index out of range. delete error")
	}
    //头节点，
	p := l.Head
	//找到要删除的节点
	for i := 0; i < index; i++ {
		p = p.Next
	}
    //这里是为了返回 要删除节点的数据（Element）
	e := p.Data

	// 先将 p [index -1] 的 Next 指向 p [index] 的 Next
	p.Pre.Next = p.Next

	// 再将 p [index + 1] 的 Pre 指向 p [index -1]
	p.Next.Pre = p.Pre

	l.Length--

	return e

}

// 查找指定位置的数据 。
func (l *List) Query(index int) Element {
	if l.IsEmpty() {
		fmt.Println("list is empty. ")

		return NoData
	}

	if index < 0 || index > l.Len() {
		return NoData
	}
    //头节点
	p := l.Head
   //循环找到节点
	for i := 0; i < index; i++ {
		p = p.Next
	}

	return p.Data
}

// 打印链表
func (l *List) Print() {
	if l.IsEmpty() {
		fmt.Println("list is empty")
	}
	//找到第一个
	p := l.Head.Next
	i := 1
	for p != nil {
		fmt.Printf("iNode %d, Data %#v\n", i, p.Data)
		i++
		p = p.Next
	}
}

func main() {
	//链表的头部节点是个空节点，也可以认为是个辅助节点。
	 //创建一个空链表，也就是空节点
	l := Create()
    //尾部插入数据
	l.Append(111)
	l.Append(222)
	l.Append(333)
	l.Append(555)
    //制定位置插入数据
	l.Insert(4, 444)
	//头部插入数据
	l.PreAppend(999)
	fmt.Println("delete ===", l.Delete(3))

	fmt.Println("query ===", l.Query(2))

	l.Print()
	fmt.Println("len ==", l.Len())
}