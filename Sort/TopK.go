package main

import (
	"fmt"
)
//借用地方写个插入排序

func Inserts(array []int)  {
	var j int
	for i:=1;i<len(array);i++{
		if array[i]<array[i-1]{
			tmp:=array[i]
			//j-1>tmp
			for j=i;j>0&&array[j-1]>tmp;j--{
				array[j] = array[j-1]
			}
			array[j] = tmp
		}
	}
}

//借用地方写快排
func Fast(array []int,i,j int){
	left,right:=i,j
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


//TopK问题

//维护小顶堆，比最小的大，放头部下沉，最终得到topK
func sink(array []int,start,length int)  {
	for{
		//2*n+1 可以画图就知道为什么了，是左边子节点
		next:=2*start+1
		if next>length{
			break
		}
		if next+1<=length&&array[next+1]<array[next]{
			next++
		}

		//上层小于下面，不用互换了
		if array[start]<=array[next]{
			break
		}
		//互换继续下沉
		array[start],array[next] = array[next],array[start]
		start = next
	}
}



//二维快排

func fast(arr [][]int,low,high int)  {
	start:=low
	end:=high
	if start>end{
		return
	}
	tmp:=arr[low]
	for start<end{
		for start<end&&arr[end][0]>=tmp[0]{
			end--
		}
		arr[start] = arr[end]
		for start<end&&arr[start][0]<=tmp[0]{
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = tmp
	fast(arr,low,start-1)
	fast(arr,start+1,high)
}

func main()  {
	array:=[]int{2,1,5,8,5,6,3,2,7}

	//tok 5
	res:=make([]int,5)
	for _,v:=range array{
		if v<=array[0]{
			continue
		}
		res[0] = v
		sink(res,0,4)
	}
	fmt.Println(res)

	Fast(array,0,len(array)-1)
	fmt.Println(array)

	st:=[][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(st)
	fast(st,0,len(st)-1)
	fmt.Println(st)


}

