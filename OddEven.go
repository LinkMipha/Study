package main

import "fmt"
//快慢指针，奇数前，偶数后
func main()  {
	nums:=[]int{1,2,3,4,5,6}
	Swap(nums)
	fmt.Println(nums)
}

func Swap (nums []int)  {
	left,right:=0,len(nums)-1
	for left<right{
		for left<right&&nums[left]&1!=0{
			left++
		}
		for left<right&&nums[right]&1!=1{
			right--
		}
		nums[left],nums[right] = nums[right],nums[left]
	}
}
