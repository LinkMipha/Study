package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//二叉搜索树
type Tree struct {
	val int
	left,right *Tree
}

func Create(root*Tree,val int)*Tree {
	root = new(Tree)
	root.val = val
	root.left = nil
	root.right = nil
	return root
}

func InsertValue(root* Tree,val int)  {
	if val>root.val{
		if root.right!=nil{
			InsertValue(root.right,val)
		}else {
			root.right=&Tree{val,nil,nil}
		}
	}else {
		if root.left!=nil{
			InsertValue(root.left,val)
		}else {
			root.left =&Tree{val,nil,nil}
		}
	}
}

func find(root*Tree,val int)bool  {
	if root==nil{
		return false
	}
	for {
		if root==nil{
			return false
		}else if val==root.val{
			return true
		}else if val>root.val{
			root = root.right
		}else {
			root = root.left
		}
	}
}

func MaxMin(root*Tree)(max,min int)  {
	for{
		if root.left!=nil{
			root  = root.left
		}else{
			min = root.val
			break
		}
	}

	for{
		if root.right!=nil{
			root = root.right
		}else{
			max = root.val
			break
		}
	}
	return min,max
}

func Print(root*Tree)[][]int {
	res := [][]int{}

	size := make([]*Tree, 0)
	size = append(size, root)

	level := 0
	for len(size) > 0 {
		res = append(res, []int{})
		length := len(size)
		for 0 < length {
			length--
			fmt.Printf("%d ", size[0].val)

			if size[0].left != nil {
				size = append(size, size[0].left)
			}
			if size[0].right != nil {
				size = append(size, size[0].right)
			}
			res[level] = append(res[level], size[0].val)
			size = size[1:]

		}
		level++
		fmt.Println()
	}
	return res
}


func main()  {
	input := bufio.NewScanner(os.Stdin)
	var root *Tree = &Tree{5,nil,nil}
	for i:=0;i<5;i++{
		input.Scan()
		value,_:=strconv.Atoi(input.Text())
		InsertValue(root,value)
	}

	//res:=Print(root)
	fmt.Println(find(root,4))
	fmt.Println(root)
	//fmt.Println(res)
}