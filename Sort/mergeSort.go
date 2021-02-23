package main

import (
	"encoding/json"
	"fmt"
)

var result []int
func mergeSort(nums []int,left,right int,res []int){
	//if len(nums)<=1{
	//	return
	//}
	//mid:=(left+right)/2//修改
	//mergeSort(nums,left,mid,res)
	//mergeSort(nums,mid+1,right,res)
	//merge(nums,left,mid,right,res)

	if left+1 < right {
		// 二分
		mid := (left+ right) / 2
		// 递归
		mergeSort(nums, left, mid, res)
		mergeSort(nums, mid+1, right, res)
		// 合并两个有序的数组
		merge(nums, left, mid, right, res)
		// 2个或者1个，只有一种条件就是左面的比右边的大才需要交换，其他条件都不需要操作
	} else if nums[left] > nums[right] {
		nums[left], nums[right]  =nums[right] , nums[left]
	}
}

func merge(data []int,start,mid,end int,buf []int)  {
	// index是合并后数组的索引，left是左半部分有序数组的索引，right是右半部分有序数组的索引
	index, left, right := start, start, mid+1
	for {
		// 假设左边的是最小的
		min := left
		if data[left] <= data[right] {
			left++
			// 如果是右半部分的最小
		} else {
			min = right
			right++
		}
		// 二者最小的元素放到缓冲中
		buf[index] = data[min]
		index++
		// 一旦有哪个半部分数组提前没有元素了就提前结束了
		if left > mid || right >= end {
			break
		}
	}
	// 把剩下的左半部分合并到缓冲中
	for i := left; i <= mid; i++ {
		buf[index] = data[i]
		index++
	}
	// 把剩下的右半部分合并到缓冲中
	for i := right; i <= end; i++ {
		buf[index] = data[i]
		index++
	}
}

func main()  {
	list := []int{22, 56, 38, 101, 1, 18, 20, 30}
	result = make([]int,len(list))
	fmt.Println("排序前", list)
	//FastSort(list, 0, len(list) - 1)
	//mergeSort(list,0,len(list)-1,result)
	//result  = Sorttwo(list)
	Me1(list)
	fmt.Println("排序后", result)
	str:="[\"basic_info\",\"sports_level\",\"sports_data\",\"medals_info\",\"group_info\"]"
	fmt.Println([]byte(str))
	var parms_list []string
	//自动解析，反序列很奇怪
	err := json.Unmarshal([]byte(str), &parms_list)
	if err!=nil{
		fmt.Println("error Unmarshal %v",err.Error())
	}
	fmt.Println(parms_list)

/*	reader:=bufio.NewReader(os.Stdin)
	n,err:=reader.ReadString('\n')
	if n==""||err!=nil{
		return
	}

	bytes, _, _ := reader.ReadLine() //处理多行时，注意一下返回值的处理
	str := string(bytes)
	fmt.Println(bytes)
	fmt.Println(str)
*/

	ch := make(chan int ,2)
	<- ch //阻塞main
}

func Sorttwo(nums []int) []int {
	if len(nums)<2{
		return nums
	}
	mid:=len(nums)/2
	left:=Sorttwo(nums[:mid])
	right:=Sorttwo(nums[mid:])
	return mergetwo(left,right)
}

func mergetwo(left,right []int)[]int  {
	res:=make([]int,0)
	m,n:=0,0
	for m<len(left)&&n<len(right){
		if left[m]<right[n]{
			res = append(res,left[m])
			m++
			continue
		}
		res = append(res,right[n])
		n++
	}
	res = append(res,left[m:]...)
	res = append(res,right[n:]...)
	return res
}



func merges1(left,right []int)[]int  {
	res:=make([]int,0)
	m,n:=0,0
	for m<len(left)&&n<len(right){
		if left[m]<right[n]{
			res = append(res,left[m])
			m++
			continue
		}
		res = append(res,right[n])
		n++
	}
	res = append(res,left[m:]...)
	res = append(res,right[n:]...)
	return res
}

func Me1(sum []int) []int {
	if len(sum)<2{
		return sum
	}
	mid:=len(sum)/2
	left:=Me1(sum[:mid])
	right:=Me1(sum[mid:])
	return merges1(left,right)
}