package main

import "fmt"

//树基本节点结构
type Node struct {
	val int
	left *Node
	right *Node
}

//树根,可用可不用
type Tree struct {
	root *Node
}


//层序构建二叉树 结合三种遍历方式，方便进行检测  
func LevelCreate(root *Node)  {


}

//创建二叉搜索树  根大于左小于右
func SearchInit(root *Node,value int)  {
	if root==nil{
		return
	}
	node:=new(Node)
	node.val  = value
	tmp:=root
	for tmp!=nil{
		//value 小于当前去左子树找
		if value<tmp.val{
			if tmp.left==nil{
				tmp.left = node
				return
			}else{
				tmp = tmp.left
			}
		}else{//否则去右子树找
			if tmp.right==nil{
				tmp.right= node
				return
			}else{
				tmp = tmp.right
			}
		}
	}
}


//前中序恢复二叉树


//中后续恢复二叉树



//中序非递归遍历 //递归简单
func MidPr(root *Node)[]int  {
	res:=make([]int,0)
	stack:=make([]*Node,0)
	tmp:=root

	for tmp!=nil||len(stack)!=0{
		for tmp!=nil{
			stack = append(stack,tmp)
			tmp = tmp.left
		}
		
		if len(stack)!=0{
			cur:=stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res,cur.val)
			if cur.right!=nil{
				tmp = cur.right
			}
		}
	}
	return res
}

//后续非递归遍历
func PostOrder(root *Node)  {

}

//层序遍历
func FourPr(root *Node)[][]int  {
	var res [][]int
	var ret []int

	queue:=make([]*Node,0)
	queue = append(queue,root)

	for len(queue)!=0{
		ret=make([]int,0)

		n:=len(queue)

		for i:=0;i<n;i++{
			tmp:=queue[i]
			queue:=queue[1:]
			ret = append(ret,tmp.val)
			if tmp.left!=nil{
				queue = append(queue,tmp.left)
			}
			if tmp.right!=nil{
				queue  = append(queue,tmp.right)
			}
		}
		res = append(res,ret)
	}
	return res
}


//二叉树之字型遍历
func OtherPr(root *Node)[][]int  {
	//两个栈，左右分别放

	var res [][]int
	var ret []int
	var sta1,sta2 []*Node
	sta1 = append(sta1,root)
	for len(sta1)!=0||len(sta2)!=0{

		ret =make([]int,0)
		for len(sta1)!=0{
			tmp:= sta1[len(sta1)-1]
			sta1 = sta1[:len(sta1)-1]
			if tmp==nil{
				continue
			}
				ret = append(ret,tmp.val)
				sta2 = append(sta2,tmp.left)
				sta2 = append(sta2,tmp.right)
		}
		if len(ret)>0{
			res = append(res,ret)
		}

		ret = make([]int,0)
		for len(sta2)!=0{
			tmp:=sta2[len(sta2)-1]
			sta2 = sta2[:len(sta2)-1]
			//记得判空
			if tmp==nil{
				continue
			}
			ret = append(ret,tmp.val)
			sta1 = append(sta1,tmp.right)
			sta1 = append(sta1,tmp.left)
		}
		if len(ret)>0{
			res = append(res,ret)
		}
	}
	return res
}

//二叉树镜像，递归，非递归方式
func MirrorRecursive(root *Node)  {
	if root==nil{
		return
	}
	//交换左右子树，下一层
	root.left,root.right  = root.right,root.left
	MirrorIteration(root.left)
	MirrorIteration(root.right)
}

//迭代镜像
//bfs队列遍历每个非叶子节点，交换它们的子节点
func MirrorIteration(root *Node) *Node {
	if root==nil{
		return nil
	}
	queue:=make([]*Node,0)
	queue  = append(queue,root)
	for len(queue)!=0{
		tmp:=queue[0]
		queue = queue[1:]
		//叶子节点不计算
		if tmp.left==nil&&tmp.right==nil{
			continue
		}
		//更换左右子树
		tmp.left,tmp.right = tmp.right,tmp.left
		if tmp.left!=nil{
			queue = append(queue,tmp.left)
		}
		if tmp.right!=nil{
			queue = append(queue,tmp.right)
		}
	}
	return root
}


//二叉树深度递归
//最大值
func Max(x,y int)int{
	if x>y{
		return x
	}else{
		return y
	}
}

//递归写起来简单
func DepthRecursive(root *Node)int  {
	if root==nil{
		return 0
	}
	left:=DepthRecursive(root.left)
	right:=DepthRecursive(root.right)
	return Max(left,right)+1
}


//二叉树深度迭代
//层序遍历记录count
func DepthIteration(root *Node) int {
	count:=0
	var res [][]int
	var ret []int

	queue:=make([]*Node,0)
	queue = append(queue,root)

	for len(queue)!=0{
		count++
		ret=make([]int,0)
		n:=len(queue)

		for n>0{
			n--
			//每次去除队列顶部元素
			tmp:=queue[0]
			queue =queue[1:]
			ret = append(ret,tmp.val)
			if tmp.left!=nil{
				queue = append(queue,tmp.left)
			}
			if tmp.right!=nil{
				queue  = append(queue,tmp.right)
			}
		}
		res = append(res,ret)
	}
	return count
}

func main()  {
	//创建树
	tree:=new(Node)
	tree.val =5
		array:=[]int{3,2,4,7,6,8}
		for _,v:=range array{
			SearchInit(tree,v)
		}
	//之字形遍历
	st:=OtherPr(tree)
	fmt.Println(st)

	//中序非递归
	mid:=MidPr(tree)
	fmt.Println(mid)

	//镜像反转
	MirrorIteration(tree)
	MirrorRecursive(tree)


	//递归求深度
	fmt.Println(DepthRecursive(tree))
	//迭代求深度
	fmt.Println(DepthIteration(tree))
}