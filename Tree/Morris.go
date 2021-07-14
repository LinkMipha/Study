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

//中序莫里斯遍历
func MidMorris(head* MorrisNode) {
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


//莫里斯后续遍历
func PostOrderMorris(head *MorrisNode) {

}


//使用数组初始化二叉树
func TreeInit(arr []int,n int,index int)*MorrisNode{
	var tree MorrisNode
	if index<n&&arr[index]!=-1{
		tree.val = arr[index]
		tree.left = TreeInit(arr,len(arr),index*2+1)
		tree.right = TreeInit(arr,len(arr),index*2+2)
	}else{
		return nil
	}
	return &tree
}


//递归先序打印
func PrePrints(head *MorrisNode){

	if head==nil{
		return
	}
	fmt.Println(head.val)
	PrePrints(head.left)
	PrePrints(head.right)
}

//递归中序
func MidPrints(head *MorrisNode)  {
	if head==nil{
		return
	}
	MidPrints(head.left)
	fmt.Println(head.val)
	MidPrints(head.right)
}

//递归后续遍历
func OldPrints(head *MorrisNode)  {
	if head==nil{
		return
	}
	OldPrints(head.left)
	OldPrints(head.right)
	fmt.Println(head.val)
}

//普通层序遍历
func FloorPrint(root *MorrisNode) []int {
	var res []int
	var queue []*MorrisNode
	queue = append(queue,root)
	for len(queue)>0{
		n:=len(queue)
		for i:=0;i<n;i++{
			tmp:=queue[0]
			queue = queue[1:]
			res = append(res,tmp.val)

			if tmp.left!=nil{
				queue = append(queue,tmp.left)
			}

			if tmp.right!=nil{
				queue = append(queue,tmp.right)
			}
		}
	}
	fmt.Println(res)
	return res
}


//非递归先序 结果对，但过程可能有问题
func NoPrePrints(head *MorrisNode)[]int{
	var res []int
	stack:=make([]*MorrisNode,0)
	stack = append(stack,head)
	for len(stack)>0{
		tmp := stack[len(stack)-1]
		res = append(res,tmp.val)
		stack = stack[:len(stack)-1]
		if tmp.right!=nil{
			stack = append(stack,tmp.right)
		}
		if tmp.left!=nil{
			stack = append(stack,tmp.left)
		}
	}
	fmt.Println(res)
	return res
}

//正确先序
func NoPrePrintsTwo(head *MorrisNode)[]int{
	var res []int
	rt:=head
	stack:=make([]*MorrisNode,0)
	for len(stack)!=0||rt!=nil{
		//左边
		for rt!=nil{
			res = append(res,rt.val)
			stack = append(stack,rt)
			rt=rt.left
		}
		//取出栈顶数据,栈顶已经放入res直接计算右侧
		tmp:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rt = tmp.right
	}
	fmt.Println(res)
	return res
}


func NoMidPrints(head *MorrisNode)[]int  {
	var res []int
	stack:=make([]*MorrisNode,0)
	rt:=head
	for len(stack)!=0||rt!=nil{
		for rt!=nil{
			stack = append(stack,rt)
			rt = rt.left
		}
		//向右
		tmp:=stack[len(stack)-1]
		res = append(res,tmp.val)
		stack = stack[:len(stack)-1]
		rt = tmp.right
	}
	fmt.Println(res)
	return res
}

//非递归后续
func NoOldPrint(head *MorrisNode) []int {
	//记录是否使用过
	var res []int
	flag:=make(map[*MorrisNode]bool)
	stack:=make([]*MorrisNode,0)
	rt:=head
	for len(stack)!=0||rt!=nil{
		for rt!=nil{
			stack = append(stack,rt)
			rt = rt.left
		}
		//取出
		tmp:=stack[len(stack)-1]
		if flag[tmp]{
			res = append(res,tmp.val)
			stack = stack[:len(stack)-1]
		}else{
			flag[tmp]=true
			rt = tmp.right
		}

	}
	fmt.Println(res)
	return res
}


func main()  {
	arr:=[]int{1,2,3,4,5}
	fmt.Println(arr[:len(arr)-1])
	point:=TreeInit(arr,len(arr),0)
	PrePrints(point)

	NoPrePrints(point)
	NoPrePrintsTwo(point)

	//mid
	NoMidPrints(point)

	//old
	NoOldPrint(point)
	//层序遍历
	//FloorPrint(point)
	//morris先序遍历
	//Morris(point)

	//midMorris
	//MidMorris(point)

}