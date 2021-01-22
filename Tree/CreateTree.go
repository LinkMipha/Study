package Tree


//树基本结构
type Node struct {
	val int
	left *Node
	right *Node
}


//创建二叉搜索树
func Search()  {

}

//不同的创建方式
//数组包含
func Create(root *Node,array []int)*Node  {

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
	var sta1,sta2 []int
	

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
func Max(x,y int)int  {
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
	return count
}

