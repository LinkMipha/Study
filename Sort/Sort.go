package main

import (
	"fmt"
	"sort"
	"sync"
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
	Name  string
}

type Te []Tem

func (t Te)Len()int  {
	return len(t)
}
func (t Te)Swap(i,j int)  {
	t[i],t[j] = t[j],t[i]
}

func (t Te)Less(i,j int) bool {
	return t[i].Score>t[j].Score
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
	fmt.Println("st",st)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b[:len(b)-1])


	t:=Te{
		{
			2,"one",
		},
		{3,"two"},
		{1,"three"},
	}

	sort.Sort(t)
	fmt.Println(t)
	//fmt.Println(b)

	fmt.Println(removeDuplicateLetters("cbacdcbc"))

	var one []int
	fmt.Println(one)
	one = make([]int,10)



	//Slice自定义排序方法
	fmt.Println("Slice自定义排序方法")
	res:= []string{"123","456","345","234","789"}
	sort.Slice(res, func(i, j int) bool {
		//前大于后则不互换
		return res[i]>res[j]
	})
	fmt.Println(res)

}

var sy sync.WaitGroup

func countDigitOne(n int) int {
	dp := make(map[int]int,0)

	for i:=1;i<=n;i++{
		sy.Add(1)
		go count(i,dp)
	}
	sy.Wait()
	res:=0
	for i:=1;i<=n;i++{
		res+=dp[i]
	}
	return dp[n]
}



func count(n int,dp map[int]int)int  {
	defer sy.Done()
	ret:=0
	for n!=0{
		s:= n%10
		if s==1{
			ret++
		}
		n = n/10
	}
	dp[n] = ret
	return ret
}

func removeDuplicateLetters(s string) string {
	stack:=make([]byte,0)
	hash:=make(map[byte]bool)
	count:=make(map[byte]int)

	for i:=0;i<len([]rune(s));i++{
		count[s[i]]++
	}

	for i:=0;i<len([]rune(s));i++{
		//栈中已经存在，不再计算
		if hash[s[i]]==true{
			continue
		}

		for len(stack)!=0&&stack[len(stack)-1]>s[i]&&count[stack[len(stack)-1]]>1{
			hash[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}

		hash[s[i]]=true
		stack = append(stack,s[i])
	}
	return string(stack)
}
