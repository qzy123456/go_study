package main

import (
	"fmt"
	"sort"
)

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// max 指向 nums1 与 nums2 合并之后的最后一个元素
	max := m + n - 1
	// 指向 num1 最后一个元素
	i := m - 1
	// 指向 num2 最后一个元素
	j := n -1
	for i >= 0 && j >= 0 {
		fmt.Println(i,j)
		// 从右向左比较值的大小
		if nums1[i] > nums2[j] {
			nums1[max] = nums1[i]
			// i 向左移动，也就是num1向左移动，往小的移动
			i--
		} else {
			nums1[max] = nums2[j]
			// j 向左移动，也就是num2向左移动，往小的方向移动
			j--
		}
		// max 向左移动
		max --
	}
	//其实下面这2个循环可以省略。因为num1的数量明显要长，所以每次都会走到 i>=0 这个里面,
	//但其实不用处理i，因为数据是有序的，j>=0也不用处理，因为 num1 > num2,所以理论上 j>= 0不会出现
	// 如果 i 越界了,num1越界了，将 nums2 剩余的元素赋值到 num1 的 [0,m] 之间
	for j >= 0 {
		nums1[max] = nums2[j]
		max--
		j--
	}
	// 如果 j 越界了，
	for i >= 0 {
		fmt.Println("我是i",i)
		nums1[max] = nums1[i]
		max--
		i--
	}
}
func merge(nums1 []int,m int,nums2 []int,n int) {
	//把nums1复制到temp中
	temp := make([]int, m)
	copy(temp, nums1)

	t, j := 0, 0 //t为temp的索引，j为nums2的索引
	for i := 0; i < len(nums1); i++ {
		//当t大于temp的长度，那就是说temp全部放进去了nums1中，那剩下的就是放nums2剩余的值了
		if t >= len(temp) {
			nums1[i] = nums2[j]
			j++
			continue
		}
		//当j大于nums2的长度的时候，那就是说明nums2全部都放进去了nums1中，那剩下的就是放temp剩余的值了
		if j >= n {
			nums1[i] = temp[t]
			t++
			continue
		}
		//比较nums2与temp对应值的大小，小的那个就放进nums1中
		if nums2[j] <= temp[t] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = temp[t]
			t++
		}
	}
}

func merge3(nums1 []int,m int,nums2 []int,n int) {
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}

func main()  {
	nums1 := []int{1,2,3}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1,m,nums2,n)
	fmt.Println(nums1)
	nums3 := []int{1,2,3,0,0,0}
	m1 := 3
	nums4 := []int{2,3,3}
	n1 := 3
	merge2(nums3,m1,nums4,n1)
	fmt.Println(nums3)
	merge3(nums3,m1,nums4,n1)
	fmt.Println(nums3)
}
