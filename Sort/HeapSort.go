package main

import (
	"fmt"
	"sort"
)

//想得到从小到大的排序结果，需要构建大顶堆，然后将堆顶最大值与最后的数据交换，
//依次进行，可以得到顺序的结果
func HeapSort(nums []int)  {

	N:=len(nums)-1
	//从底部到顶部构建大顶堆，最后一个非叶子节点开始
	for i:=N/2;i>=0;i--{
		sink(nums,i,N)
	}

	//将堆顶值和末尾交换，重新调整堆
	for i:=N;i>=0;i--{
		wap(nums,0,i)
		sink(nums,0,i-1)//交换之后，数组最后一位不算在堆内，需要减1操作
	}

	//不同写法，结果一样
	/*for N>=0{
		wap(nums,0,N)
		N--
		sink(nums,0,N)
	}*/
}

func sink(nums []int,k,N int)  {
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



func Testsink(sum []int,k,len int)  {
	for{
		i:=k*2+1
		if i>len{
			break
		}
		if i+1<=len&&sum[i+1]>sum[i]{
			i++
		}
		if sum[k]>=sum[i]{
			break
		}
		sum[k],sum[i] = sum[i],sum[k]
		k = i
	}
}

func Testheap(nums []int)  {
	length:=len(nums)-1

	for i:=length/2;i>=0;i--{
		Testsink(nums,i,length)
	}

	for i:=length;i>=0;i--{
		nums[0],nums[i] = nums[i],nums[0]
		Testsink(nums,0,i-1)
	}
}

func main()  {
	//s := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	s :=[]int{2,1}
	fmt.Println(s)
	//HeapSort(s)
	//Testheap(s)
	sort.Ints(s)
	fmt.Println(s)
	defer fmt.Println()

	//var st  = 1
	//st = ^st+1
	//fmt.Println(st)
	//fmt.Println(^9)
	//
	//a := "aaa"
	//ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	//b := *(*[]byte)(unsafe.Pointer(&ssh))
	//fmt.Printf("%s", b)


	//var test *Test = new(Test)
	//fmt.Println(test)
	//
	//a := unsafe.Pointer(uintptr(unsafe.Pointer(test))+unsafe.Offsetof(test.a))
	//*((*int)(a)) = 21
	//fmt.Println(test)


}