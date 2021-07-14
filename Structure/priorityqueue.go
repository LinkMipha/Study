package main

import "fmt"

//优先队列，参考源码

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0{
		return nil
	}

	var  heap []*ListNode
	for i:=0;i<len(lists);i++{
		if lists[i]!=nil{
			heap = push(heap,lists[i])
		}
	}
	sink(heap,0)//形成堆
	dummy:=new(ListNode)
	cur:=dummy
	for len(heap)!=0{
		var tmp *ListNode
		heap,tmp = pop(heap)//取出堆顶

		cur.Next = tmp
		cur = cur.Next
		if tmp.Next!=nil{  //非空放入下一个
			heap = push(heap,tmp.Next)
		}
	}
	return dummy.Next
}

func sink(arr []*ListNode,start int)  {
	length:=len(arr)-1
	for{
		next:=start*2+1
		if next>length{
			return
		}
		if next+1<=length&&arr[next+1].Val<arr[next].Val{
			next++
		}
		if arr[start].Val<arr[next].Val{
			break
		}
		arr[start],arr[next] = arr[next],arr[start]
		start = next
	}
}

//节点上升只需要找父节点
func swim(arr []*ListNode,start int){
	for {
		parent:=(start-1)/2
		if start==parent||arr[parent].Val<=arr[start].Val{
			break
		}
		arr[parent],arr[start] = arr[start],arr[parent]
		start = parent
	}
}


//insert
func push(arr []*ListNode,root *ListNode)[]*ListNode  {
	//上升节点
	arr = append(arr,root)
	swim(arr,len(arr)-1)
	return arr
}

func pop(arr []*ListNode)([]*ListNode,*ListNode)  {
	tmp:=arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	sink(arr,0)
	return arr,tmp
}

func Creat(arr []int)*ListNode  {
	root:=new(ListNode)
	cur:=root
	for i:=0;i<len(arr);i++{
		tmp:=new(ListNode)
		tmp.Val = arr[i]
		cur.Next = tmp
		cur = cur.Next
	}
	return root.Next
}

func main()  {
	one:=Creat([]int{-5,1,1,3,4})
	two:=Creat([]int{-2})
	three:=Creat([]int{0,1,3})
	four:=Creat([]int{2})
	var lists = []*ListNode{one,two,three,four}
	st:=mergeKLists(lists)
	fmt.Println(st)

}