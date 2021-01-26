package main

import (
	"fmt"
	"sync"
)

var one= make(chan int,1)
var two = make(chan int,1)
var group sync.WaitGroup

//可以使用chan代替group
var end = make(chan int,1)

func PrintA()  {
	//defer group.Done()
	//两个chan控制打印
	for i:=0;i<5;i++{
		<-one
		fmt.Println("A")
		two<-1
	}
}

func PrIntB()  {
	//defer group.Done()
	for i:=0;i<5;i++{
		<-two
		fmt.Println("B")
		one<-1
	}
	end<-1
}

func Three(c1 chan bool,c2 chan bool,str string)  {
	//value 和 真假
	if _,ok:=<-c1;ok{
		fmt.Println(str)
		c2<-true
	}
}

func main()  {
	//启动信号
	one<-1
	//group.Add(2)
	go PrintA()
	go PrIntB()
	//group.Wait()
	<-end
}
