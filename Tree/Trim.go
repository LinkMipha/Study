package main

import "fmt"

//前缀树/字典树

type Node struct {
	Next [26]*Node
	Sum int
	isEnd bool //增加bool判断是否到达结尾
}
var trim [1000][128]int
var TrimSum []int
var trimCount int

//字典树

func Insert(str string,root *Node)  {
	head:=root
	for i:=0;i<len(str);i++{
		d:=str[i]-'a'
		if head.Next[d]==nil{
			node:=new(Node)
			head.Next[d]=node
			head.Sum++
		}
		head = head.Next[d]
	}
}



func Search(str string,root *Node) bool {
	head:=root
	for i:=0;i<len(str);i++{
		d:=str[i]-'a'
		if head.Next[d]==nil {
			return false
		}else{
			head = head.Next[d]
		}
	}
	if head.Sum==0{
		return true
	}else{
		return false
	}

}

func StartsWith(str string,root *Node)bool  {
	head:=root
	for i:=0;i<len(str);i++{
		d:=str[i]-'a'
		if head.Next[d]==nil{
			return false
		} else{
			head = head.Next[d]
		}
	}
	return true
}

func main()  {
	root:=new(Node)
	Insert("apple",root)
	fmt.Println(Search("apple",root))
	fmt.Println(Search("app",root))
	fmt.Print(StartsWith("app",root))
}