package main

import "fmt"

func Sink(nums[]int,i int,length int)  {
	for {
		start:=i*2+1
		if start>length{
			break
		}
		if start+1<=length&&nums[start+1]>nums[start]{
			start++
		}
		if nums[i]>nums[start]{
			break
		}
		Wap(nums,i,start)
		i = start
	}
}

func Wap(nums[]int,a,b int)  {
	nums[a],nums[b] = nums[b],nums[a]
}

func Init(nums []int)  {
	length:=len(nums)-1
	//构建堆
	for i:=length/2;i>=0;i--{
		Sink(nums,i,length)
	}

	//交换排序
	for i:=length;i>=0;i--{
		Wap(nums,0,i)
		Sink(nums,0,i-1)
	}
}

func main()  {
	//s := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	arr:=[]int{2,1,4,4,5,4,3,2,7}
	fmt.Println(arr)
	Init(arr)
	fmt.Println(arr)
}