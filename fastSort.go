package main

import "fmt"

func FastSort(st []int,start,end int)  {
	if start>=end{
		return
	}
	low:=start
	high:=end

	mid:=st[start]
	for low<high{
		for low<high&&st[high]>=mid{
			high--
		}
		st[low],st[high] = st[high],st[low]
		for low<high&&st[low]<=mid{
			low++
		}
		st[low],st[high] = st[high],st[low]
	}
	st[low] = mid
	FastSort(st,start,low-1)
	FastSort(st,low+1,end)
}

func TestFastSort(sum []int,start,end int){
	if start>=end{
		return
	}
	low:=start
	high:=end
	mid:=sum[start]
	for low<high{
		for low<high&&sum[high]>=mid{
			high--
		}
		sum[low],sum[high] = sum[high],sum[low]
		for low<high&&sum[low]<=mid{
			low++
		}
		sum[low],sum[high] = sum[high],sum[low]
	}
	sum[low] = mid
	TestFastSort(sum,start,low-1)
	TestFastSort(sum,low+1,end)
}


func main()  {
	list := []int{22, 56, 38, 101, 1, 18, 20, 30}
	fmt.Println("排序前", list)
	TestFastSort(list,0,len(list)-1)
	//fast(list,0,len(list)-1)
	fmt.Println("排序后", list)
}

func fast(st []int,l,r int)  {
	if l>=r{
		return
	}
	p:=SortTwo(st,l,r)
	fast(st,l,p-1)
	fast(st,p+1,r)
}

func SortTwo(st []int,l,r int)int {
	v:=st[l]
	j:=l
	//找右边第小于初始值的数据放在前面
	for i:=j+1;i<=r;i++{
		if st[i]<v{
			j++
			st[i],st[j] = st[j],st[i]
		}
	}
	//最终st[j]位置一定不大于 最初的标准值，将标准值放中间。
	st[j],st[l] = v,st[j]
	return j
}

