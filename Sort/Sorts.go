package main

import (
	"fmt"
)


//冒泡排序
//稳定排序，平均复杂O(n^2) 最优 O(n)

//初始版本
func Bubble(array []int)  {
	//次数小于长度-1，因为内层循环会+1
	for i:=0;i<len(array)-1;i++{
		for j:=0;j<len(array)-i-1;j++{
			if array[j]>array[j+1]{
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
	}
}

//优化方式1,增加flag判断
func BubbleOne(array []int)  {
	flag:=false
	for i:=0;i<len(array)-1;i++{
		flag = false

		for j:=0;j<len(array)-i-1;j++{
			if array[j]>array[j+1]{
				array[j],array[j+1] = array[j+1],array[j]
				flag = true
			}
		}
		//未执行则已经有序
		if !flag{
			return
		}
	}
}

//优化方式2，在1的基础上
//增加position记录最后交换的位置j，下次不再判断后面的数字
func Bubbletwo(array []int) {
	//加入flag,判断是否提前结束
	flag:=false
	length:=len(array)-1
	//记录最后一次交换的位置，后面没有交换，为有序
	position:=0
	for i:=0;i<len(array)-1;i++{
		flag = false
		for j:=0;j<length;j++ {
			if array[j]>array[j+1]{
				array[j],array[j+1] = array[j+1],array[j]
				flag = true
				position = j//记录最后一次交换次数
			}
			fmt.Printf("外层%d，内层%d\n",i,j)
		}
		if !flag{
			break
		}
		length  = position
	}
}


//选择排序，选择最小和开始交换
//好坏平均复杂度 O(n^2) 不稳定
func SelectSort(array []int)  {
	index:=0
	for i:=0;i<len(array)-1;i++{
		index =  i
		for j:=i+1;j<len(array);j++{
			//比最小的小，则更新最小index
			if array[index]>array[j]{
				index = j
			}
		}
		//交换
		array[i],array[index]=array[index],array[i]
	}
}

//插入排序  稳定排序  时间O(n^2)
func Inserts(nums []int,n int)  {
	var j int
	for i:=1;i<n;i++{
		if nums[i]<nums[i-1]{
			//保存临时数据，数组数据后移
			tmp:=nums[i]
			//注意边界 j>0
			for j=i;j>0&&nums[j-1]>tmp;j--{
				nums[j] = nums[j-1]
			}
			nums[j] = tmp
		}
	}
}


//快排	不稳定排序  O(nlogn)  还有多路快排优化
//j = len(array)-1
func Fast(array []int,i,j int)  {
	left:=i
	right:=j
	//超出范围 退出，递归停止条件
	if left>=right{
		return
	}

	tmp:=array[left]
	for left<right{
		for left<right&&array[right]>=tmp{
			right--
		}
		array[left] = array[right]
		for left<right&&array[left]<=tmp{
			left++
		}
		array[right] = array[left]
	}
	count++
	fmt.Println("time",count)
	array[left] = tmp
	Fast(array,i,left-1)
	Fast(array,left+1,j)
}

//归并   稳定排序  时间O(nlogn)  空间O(n) 有多路归并，
// 归并可以结合很多其他知识 链表，树等等,分治思想很重要（至少面试会问）
// 海量数据排序可以使用多路归并
func MergeSort(array []int)[]int  {
	//单个数据时返回，递归停止条件
	if len(array)<2{
		return array
	}

	//递归分割左右，想好递归停止的条件，分治
	mid:=len(array)/2
	left:=MergeSort(array[:mid])
	right:=MergeSort(array[mid:])
	return merge(left,right)
}

//合并两个有序链表类似的做法
func merge(left,right []int) []int {
	//额外的返回空间
	res:=make([]int,0)
	i,j:=0,0

	for i<len(left)&&j<len(right){
		//<=是稳定的关键点
		if left[i]<=right[j]{
			res = append(res,left[i])
			i++
		}else{
			res = append(res,right[j])
			j++
		}
	}

	//这里偷懒了，可以自己写详细些
	/*if i==len(left){
		res = append(res,right[j:]...)
	}else{
		res = append(res,left[i:]...)
	}*/

	//详细写法
	//left right 其中一个没有结束
	if i<len(left){
		for i<len(left){
			res = append(res,left[i])
			i++
		}
	}else{
		for j<len(right){
			res = append(res,right[j])
			j++
		}
	}
	return res
}

//希尔排序
//间隔插入排序



//堆排序   不稳定  时间复杂O(nlogn)
func HeapSort (array []int)  {
	//建堆，从后向前建立堆
	for i:=len(array)/2-1;i>=0;i--{
		sink(array,i,len(array)-1)
	}

	//互换并继续维护堆
	for j:=len(array)-1;j>=0;j--{
		array[0],array[j] = array[j],array[0]
		sink(array,0,j-1)
	}
}

//下沉，堆排序最关键堆核心
//想得到从小到大的数据，需要大顶堆，最大放在结尾，最后是小到大顺序
func sink(array []int,start ,length int)  {
	for{
		next:=2*start+1
		//超出范围 退出循环
		if next>length{
			break
		}
		//找最大的值和上层互换  <=length很关键
		if next+1<=length&&array[next+1]>array[next]{
			next++
		}
		//上层已经大于下面，不用下沉，直接退出循环
		if array[start]>array[next]{
			break

		}
		//互换后继续寻找，可能继续下沉
		array[start],array[next] = array[next],array[start]
		start = next
	}
}

func HeapSorts(nums []int)  {

	N:=len(nums)-1
	//从底部到顶部构建大顶堆，最后一个非叶子节点开始
	for i:=N/2;i>=0;i--{
		sinks(nums,i,N)
	}

	//将堆顶值和末尾交换，重新调整堆
	for i:=N;i>=0;i--{
		wap(nums,0,i)
		sinks(nums,0,i-1)//交换之后，数组最后一位不算在堆内，需要减1操作
	}

	//不同写法，结果一样
	/*for N>=0{
	    wap(nums,0,N)
	    N--
	    sink(nums,0,N)
	}*/
}

func sinks(nums []int,k,N int)  {
	for{
		i:=2*k+1
		if i>N{
			break
		}
		//找左右子节点最大值
		if i+1<=N&&nums[i+1]>nums[i]{
			i++
		}
		//已经大于最大值，不需要再交换
		if nums[k]>=nums[i]{
			break
		}
		wap(nums,k,i)
		k = i //继续向上调整
	}
}

func wap(nums []int,x,y int){
	nums[x],nums[y] = nums[y],nums[x]
}


var count int
//三路快排，分三部分，1小于tmp部分，等于tmp部分，大于tmp部分
func FastThree(array []int,low,high int){
	i:=low
	j:=low
	tmp:=array[low]
	k:=high
	if low>=high{
		return
	}
	//此处为k
	for i<=k{
		if array[i]<tmp{
			//小于部分和前面等于的最左边互换，然后当前位置不比较，++
			array[i],array[j] = array[j],array[i]
			i++
			j++
		}else if array[i]>tmp{
			//大于的放在右边，然后i不加，因为更换位置后此处的大小未知
			array[k],array[i] = array[i],array[k]
			k--
		}else{
			i++
		}
	}
	count++
	fmt.Println("time",count)
	FastThree(array,low,j-1)
	FastThree(array,k+1,high)
}


func main()  {
	array:=[]int{22, 56, 38, 101, 1, 18, 20, 30}
//	arr:=[]int{1,2,2,3,5,5,6,7,8}
	//arra:=[]int{2,1,5,8}
	//insertSort(array,9)
	//Fast(array,0,8)
    //CreateHeap(array)
	//Inserts(array,9)

	//SelectSort(array)
	fmt.Println(array)

	//res:=MergeSort(array)
	//FastThree(array,0,len(array)-1)
	HeapSort(array)
	fmt.Println(array)

	//fmt.Println(array[:len(array)])

	one:="我来测试"
	fmt.Println(len(one))
	fmt.Println(len([]rune(one)))

}