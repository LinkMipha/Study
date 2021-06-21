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
		//找左子树最右
		if cur.left!=nil{
			var pre *MorrisNode = cur
			for pre.right!=nil&&pre.right!=cur{
				pre = pre.right
			}
			if pre.right==nil{
				fmt.Println(cur.val)
				pre.right=cur
				cur = cur.left
			}else{
				pre.right=nil
				cur = pre.right
			}

		}else{
			fmt.Println(cur.val)
			cur = cur.right
		}
	}
}


//初始化
func TreeInit(arr []int,len int,index int)*MorrisNode{
	var tree MorrisNode
	tree.left = nil
	tree.right = nil
	if index<len&&arr[index]!=-1{
		tree.val = arr[index]
		tree.left = TreeInit(arr,len,index*2+1)
		tree.right = TreeInit(arr,len,index*2+2)
	}
	if tree.left==nil&&tree.right==nil{
		return nil
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
	Prints(point)


	//morris先序遍历
	//Morris(point)
}