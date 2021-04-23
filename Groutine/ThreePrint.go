package main

import (
	"fmt"
	"time"
)

//两个channel交替打印切片数据
var arr = []int{1,2,3,4,5}

func g1(c1 chan bool)  {
	for i:=0;i<len(arr);i++{
		c1<-true
		if i%2==0{
			fmt.Println(arr[i])
		}
	}
}

func g2(c2 chan bool)  {
	for i:=0;i<len(arr);i++{
		<-c2
		if i%2==1{
			fmt.Println(arr[i])
		}
	}
}

//两个groutine交替打印数字






func main()  {
	gr:=make(chan bool)
	go g1(gr)
	go g2(gr)
	time.Sleep(time.Duration(1)*time.Second)
}