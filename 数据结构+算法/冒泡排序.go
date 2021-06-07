package main

import "fmt"

func BubbleSort(arr []int)  {
	for i:=0;i<len(arr) ;i++  {
		for j:=0;j<len(arr) -1-i ;j++  {
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

func main()  {
	sliec := []int{23,14,21,211}
	BubbleSort(sliec)
	fmt.Println(sliec)
}