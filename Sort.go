package main

import (
	"fmt"
	"sort"
)

type student struct {
	name string
	distance int
}

type Str []student

func (st Str)Len()int  {
	return len(st)
}
func (st Str)Swap(x,y int)  {
	st[x],st[y] = st[y],st[x]
}

func (st Str)Less(x,y int) bool {
	return st[x].distance>st[y].distance
}

type Tem struct {
	Score int
	Name  stirng
}

func Type string ()  {
	
}



func main()  {
	a:=Str{
		{name: "one",distance: 12},
		{
			name: "two",
			distance: 14,
		},
		{
			name: "three",
			distance: 13,
		},
		{
			name:"four",
			distance: 16,
		},
	}
	sort.Sort(a)
	fmt.Println(a)
	b := []int{4,3,2,1,5,9,8,7,6}
	st:=sort.IntSlice(b)
	fmt.Println(st)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b[:len(b)-1])
	
	
	//fmt.Println(b)
}

