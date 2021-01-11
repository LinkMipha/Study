package main

import "fmt"

//calculate next array
func Next(p string,n int)[]int {
	next:=make([]int,n)
	next[0] = -1
	j:=-1
	i:=0
	for i<len(p)-1{
		if j==-1||p[i]==p[j]{
			i++
			j++
			next[i] = j
		}else{
			j = next[j]
		}
	}
	return next
}

//第二种求解next数组方法
//next 为匹配位置-1
func Ne(str string) []int {
	length:=len(str)
	next:=make([]int,length)
//dp 思想  每次从已知状态转换到新到状态 不匹配则转换到next位置
	next[0] = -1
	for i:=1;i<length;i++{
		j:=next[i-1]

		for str[j+1]!=str[i]&&j>=0{
			j = next[j]
		}
		if str[j+1]==str[i]{
			next[i] = j+1
		}else{
			next[i] = -1
		}
	}
	return next
}

//obtain kmp length
func Get(s string,p string)[]int{
	array:=make([]int,0)
	i:=0
	j:=0
	for i<len(s){

		for i<len(s)&&j<len(p){

			if j==-1||s[i]==p[j]{
				i++
				j++
			}else{
				j = Next(p,len(p))[j]
			}
		}

		if j==len(p){
			array = append(array,i-j)
			j = Next(p,len(p))[j]
		}
	}

	return  array
}

func main()  {
	str := "ABABABC"
	p:="ABA"
	fmt.Println(Next(p,len(p)))
	fmt.Println(Get(str,p))
}