package main

import (
	"fmt"
)

type MorrisNode struct {
	val int
	left *MorrisNode
	right *MorrisNode
}

//莫里斯先序遍历
func Morris(head *MorrisNode)  {
	var cur *MorrisNode = head
	for cur!=nil{
		if cur.left==nil{
			fmt.Println(cur.val)
			cur = cur.right
		}else{
			//找左子树最右
			var pre = cur.left
			for pre.right!=nil&&pre.right!=cur{
				pre = pre.right
			}
			if pre.right==nil{
				fmt.Println(cur.val)
				pre.right=cur
				cur = cur.left
			}else{
				pre.right=nil
				cur = cur.right
			}

		}
	}
}

func MidMorris(head* MorrisNode)  {
	var cur = head
	for cur!=nil{
		if cur.left==nil{
			fmt.Println(cur.val)
			cur =cur.right
		}else{
			//左最右
			var pre = cur.left
			for pre.right!=nil&&pre.right!=cur{
				pre = pre.right
			}
			if pre.right==nil {
				pre.right = cur
				cur = cur.left
			}else {
				fmt.Println(cur.val)
				pre.right = nil
				cur = cur.right
			}
		}
	}
}

func PostOrderMorris(head *MorrisNode) {

}


//初始化
func TreeInit(arr []int,n int,index int)*MorrisNode{
	var tree MorrisNode
	if index<n&&arr[index]!=-1{
		tree.val = arr[index]
		tree.left = TreeInit(arr,len(arr),index*2+1)
		tree.right = TreeInit(arr,len(arr),index*2+2)
	}
	return &tree
}


//合理
func Prints(head *MorrisNode)  {
	if head==nil{
		return
	}
	fmt.Println(head.val)
	Prints(head.left)
	Prints(head.right)
}

func main()  {
	arr:=[]int{1,2,3,4,5}
	point:=TreeInit(arr,len(arr),0)
	//Prints(point)

	//morris先序遍历
	//Morris(point)

	//midMorris
	MidMorris(point)
}