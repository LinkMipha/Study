package main
import "fmt"

//链表节点
type Node struct{
	Next *Node
	value int
}

//构建
func Create(array []int)*Node{
	root :=new (Node)
	cur:=root
	for i:=0;i<len(array);i++{
		tmp:=new(Node)
		tmp.value = array[i]
		cur.Next = tmp
		cur = cur.Next
	}
	return root.Next
}

//输出
func Print(root *Node){
	for root!=nil{
		fmt.Printf("%v",root.value)
		root = root.Next
	}
}


//反转
func Reverse(root *Node)*Node{

	if root==nil||root.Next==nil{
		return root
	}
	//pre初始为nil
	var pre *Node
	cur:=root
	for cur!=nil{
		tmp:=cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}


//找到倒数第k个节点
func FindK(root *Node,k int)int{
	//双指针先到第k
	fast:=root
	for k>0&&fast!=nil{
		fast = fast.Next
		k--
	}
	//超出长度
	if k>0{
		return -1
	}
	for fast!=nil{
		fast = fast.Next
		root = root.Next
	}
	return root.value
}

//删除倒数第k个节点,和倒数第k有区别
func DeleteK(root *Node,k int){
	dummy:=new(Node)
	dummy.Next = root
	pre:=dummy
	//双指针先到第k
	fast:=root
	for k>0&&fast!=nil{
		fast = fast.Next
		k--
	}
	//超出长度
	if k>0{
		return
	}
	for fast!=nil{
		fast = fast.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return
}


//原地删除链表节点
func DeleteOne(node *Node)  {
	//改值改节点
	node.value = node.Next.value
	node.Next = node.Next.Next
}


//从m-n反转链表
func ReverseMtoN(root *Node,m,n int)*Node{
	//头插法反转链表
	//先找到m前一个节点,头插入n-m次（一共n-m+1个，所以是n-m次）
	dummy:=new(Node)
	dummy.Next = root
	pre:=dummy
	for i:=0;i<m-1&&pre!=nil;i++{
		pre = pre.Next
	}
	//越界
	if pre==nil{
		return root
	}
	nextHead:=pre.Next
	move:=nextHead.Next
	for i:=0;i<n-m;i++{
		nextHead.Next = move.Next
		move.Next = pre.Next
		pre.Next = move
		move = nextHead.Next
	}
	return dummy.Next
}


//k个一组反转链表
func ReverseK(root *Node,k int)*Node{
	length:=0
	cur:=root
	for cur!=nil{
		length++
		cur = cur.Next
	}
	//翻转time次每次k-1个
	time:=length/k
	//返回最后结果
	dummy:=new(Node)
	dummy.Next = root
	pre:=dummy
	var nextHead,move *Node
	for i:=0;i<time;i++{
		//nextHead为开始，move为第二个，每次move放nextHead前面
		nextHead = pre.Next
		move = nextHead.Next
		for j:=0;j<k-1;j++{
			nextHead.Next = move.Next
			move.Next= pre.Next
			pre.Next = move
			move = nextHead.Next
		}
		//结束时更改pre
		pre = nextHead
	}
	return dummy.Next
}

func main() {
	arr:=[]int{1,2,3,4,5,6,7}
	//创建链表
	tmp:=Create(arr)

	//反转链表
	//root = Reverse(tmp)

	//查找倒数第k个
	fmt.Println(FindK(tmp,3))

	//删除第三个
	//DeleteK(tmp,3)

	//原地删除
	//DeleteOne(tmp)


	//m-n反转
	//ReverseMtoN(tmp,2,4)

	//k个一组翻转链表（多余的不反转）
	tmp = ReverseK(tmp,3)

	//输出链表
	Print(tmp)
}