package main

import "fmt"

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
//你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。
//返回滑动窗口中的最大值所构成的数组。
//超时
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums) //数组长度
	if length == 1{
		return []int{nums[0]}
	}
	index := 0      //起始下标
	ret := make([]int, 0) //返回值
	for index < length {
		m := nums[index]
		//不够分组了
		if index > length - k {
			break
		}
		//第二个开始比较，k个为一组。找出最大值
		for j := index + 1; j < index + k; j++ {
			if m < nums[j] {
				m = nums[j]
			}
		}
		ret = append(ret,m)
		index++
	}
	return ret
}

// 解法二 双端队列 Deque
//最优的解法是用双端队列，队列的一边永远都存的是窗口的最大值，队列的另外一个边存的是比最大值小的值。
//队列中最大值左边的所有值都出队。在保证了双端队列的一边即是最大值以后，时间复杂度是 O(n)，空间复杂度是 O(K)
//利用一个双端队列来保存当前窗口中最大那个数在数组里的下标，双端队列新的头就是当前窗口中最大的那个数。通过该下标，
//可以很快地知道新的窗口是否仍包含原来那个最大的数。如果不再包含，我们就把旧的数从双端队列的头删除。
//因为双端队列能让上面的这两种操作都能在 O(1) 的时间里完成，所以整个算法的复杂度能控制在 O(n)。
//1、初始化窗口 k=3，包含 1，3，-1，把 1 的下标压入双端队列的尾部；
//2、把 3 和双端队列的队尾的数据逐个比较，3 >1，把 1 的下标弹出，把 3 的下标压入队尾；
//3、-1<3，-1 压入双端队列队尾保留到下一窗口进行比较；
//4、3 为当前窗口的最大值；
//5、窗口移动，-3 与队尾数据逐个比较，-3<-1，-3 压入双端队列队尾保留；
//6、3 为当前窗口的最大值；
//7、窗口继续移动，5>-3，-3 从双端队列队尾弹出；
//8、5>-1，-1 从队尾弹出；
//9、3 超出当前窗口，从队列头部弹出；
//10、5 压入队列头部，成为当前窗口最大值；
//11、继续移动窗口，操作与上述同理。
//窗口最大值只需读取双端队列头部元素。
//要分清头和尾
//演绎维护单调双端队列的4个步骤
//一四个步骤的目的，保证滑动窗口头部是最大值
//(头，尾，尾，头)
//第一步，头部出队，清理超范围（头部出队只出一个元素）
//第二步，篮球队清理，移除尾部，在当前值前面的还小于当前值的元素
//第三步，尾部入队，尾部范围正确
//第四步，返回头部一一当前窗口最大值
//
//比如【1 3 -1 -3 5 3 6 7】
//先进入1，扫描3，发现3>1，然后1出队，3进队。-1<3，-1进队，最大值为3；
//-3<-1, -3进队，返回头部，最大值为3；
//清除头部3，5>-3, -3出队， 5>-1, -1出队，5进队，返回头部，最大值为5；
//清除头部空格，3<5, 3进队，返回头部最大值为5；
//清除头部空格，扫描6，6>3, 3出队，6>5， 5出队，返回头部最大值为6；
//清除头部空格，7>6，6出队，7入队，返回头部最大值7；
//最后结果为335567；
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k)
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if i >= k && window[0] <= i-k {     // 如果最左边的索引在窗口外，将其删除
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v {  // 维护窗口
			window = window[ : len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]]) // 最左边是nums中最大值的索引
		}
	}
	return result
}
//解法3 单调队列
//单调递增栈：从栈顶到栈底，依次递增
//单调递减栈：从栈顶到栈底，依次减小
func maxSlidingWindow3(nums []int, k int) []int {
	var q []int
	push := func(i int) {
		//维护递减队列
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
    //nums ==>  1,3,-1,-3,5,3,6,7, k==> 3
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	//fmt.Println(q) //1 2
	for i := k; i < n; i++ {
		push(i)
		//fmt.Println(nums[q[0]],q[0],i,k,i-k)
		//3 1 3 3 0
		//5 4 4 3 1
		//5 4 5 3 2
		//6 6 6 3 3
		//7 7 7 3 4
		//k个一区间
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	//[3 3 5 5 6 7]
	return ans
}

func main() {
	//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3 输出: [3,3,5,5,6,7] 解释:
   //[1 3 -1] -3 5 3 6 7  输出 3
	//
	//1 [3 -1 -3] 5 3 6 7  输出 3
	//
	//1 3 [-1 -3 5] 3 6 7  输出 5
	//
	//1 3 -1 [-3 5 3] 6 7  输出 5
	//
	//1 3 -1 -3 [5 3 6] 7  输出 6
	//
	//1 3 -1 -3 5 [3 6 7]  输出 7
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	//fmt.Println(maxSlidingWindow(nums,k))
	//fmt.Println(maxSlidingWindow2(nums,k))
	fmt.Println(maxSlidingWindow3(nums,k))
}
