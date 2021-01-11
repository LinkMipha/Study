package main

import "fmt"


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
	if left>right{
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
	length:=len(array)-1
	//建堆，从后向前建立堆
	for i:=len(array)/2;i>=0;i--{
		sink(array,i,length)
	}

	for i:=length;i>=0;i--{
		array[0],array[i] = array[i],array[0]
		sink(array,0,i-1)
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

		//找最大的值和上层互换
		if next+1<length&&array[next+1]>array[next]{
			next++
		}

		//上层已经大于下面，不用下沉，直接退出循环
		if array[start]>=array[next]{
			break
		}
		//互换后继续寻找，可能继续下沉
		array[start],array[next] = array[next],array[start]

		start = next
	}
}


func main()  {
	array:=[]int{2,1,5,8,5,6,3,2,7}
	//arra:=[]int{2,1,5,8}
	//insertSort(array,9)
	//Fast(array,0,8)
    //CreateHeap(array)
	//Inserts(array,9)

	//res:=MergeSort(array)
	HeapSort(array)
	fmt.Println(array)


	fmt.Println(array[:len(array)])

	one:="我来测试"
	fmt.Println(len(one))
	fmt.Println(len([]rune(one)))

}