package main

import "fmt"

type NodeMorris struct {
	val int
	left *NodeMorris
	right *NodeMorris
}

func Morris(head* NodeMorris)  {
	//cur
	cur:=head
	for cur!=nil{
		//寻找左侧最右
		if cur.left!=nil{
			var pre = cur.left
			for pre.right!=nil&&pre.right!=cur{
				pre = pre.right
			}
			// one
			if pre.right==nil{
				fmt.Println(cur.val)
				pre.right = cur
				cur = cur.left
			}

			if pre.right==cur{
				pre.right=nil
				cur = cur.right
			}

		}else{
			fmt.Println(cur.val)
			cur =cur.right
		}
	}

}

func TreeInit(arr []int,index int,len int)*NodeMorris  {
	var node NodeMorris
	if index<len&&arr[index]!=-1{
		node.val = arr[index]
		node.left = TreeInit(arr,index*2+1,len)
		node.right = TreeInit(arr,index*2+2,len)
	}
	return &node
}


func TreePrint(head * NodeMorris){
	if head==nil{
		return
	}
	fmt.Println(head.val)
	TreePrint(head.left)
	TreePrint(head.right)
}

func main()  {

	for i:=0;i<10;i++{
		
	}
	var q [][]int
	for a,b:=2,1;a>0&&b>0;b=b-2{
		q = append(q,[]int{1,a++})
	}

	arr:=[]int{1,2,3,4,5}
	head:=TreeInit(arr,0,len(arr))
	TreePrint(head)
    //Morris(head)

    one:=[]int{1,2,3,4,5}
    fmt.Println(one[1:len(one)-1])
    copy()
}