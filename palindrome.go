package main

import "fmt"

//判断回文字符串
func find(s string)bool  {
	var ret []rune = []rune(s)
	length:=len(ret)
	for i:=0;i<length/2;i++{
		if s[i]!=s[length-i-1]{
			return false
		}
	}
	return true
}

//第二种方式判断
func isPalindrome(s string)bool {
	length := len([]rune(s))
	left,right:=0,length-1
	for left<right{
		if s[left]!=s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

//动态规划判断回文

func dp(s string)bool  {
	return  true
}

func main()  {
	st:="a"
	fmt.Println(isPalindrome(st))
	ret:="hello"
	fmt.Println(string(ret[1:2]))
}