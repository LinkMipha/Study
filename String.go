package main

import (
	"fmt"
	"strings"
)

func main()  {
	//stting包的练习
	str :=" he ha he st hs "
	//大小写
	sta:=strings.ToUpper(str)
	statwo :=strings.ToLower(str)
	sl:=strings.Split(str," ")

	trim:=strings.TrimSpace(str)
	fmt.Println("len",len(sl),sl)

	sp:=strings.Join(sl,",")
	fmt.Println("join",sp)

	fmt.Println("up,low",sta,statwo)
	fmt.Println("trim",trim,"a")


	fmt.Println(strings.Repeat("hello",5))

	//从前开始去除包含在后续字符串中的字符，直到新的字符没有为止
	//strings.TrimRight()同理
	fmt.Println(strings.TrimLeft("bababcababab","ba"))

	fmt.Println(strings.TrimRight("cyeamblog.go", ".go"))
	fmt.Println(strings.TrimSuffix("cyeamblog.go", ".go"))

	//将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
	fmt.Println(strings.Fields(" he ha he st hs "))
}