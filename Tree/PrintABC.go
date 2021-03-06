package main

import (
	"fmt"
	"sync"
	"time"
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
	//defer group.Done()
	//value 和 真假
	if _,ok:=<-c1;ok{
		fmt.Println(str)
		c2<-true
	}
}

func main()  {
	//启动信号
	//one<-1
	//group.Add(2)
	go PrintA()
	go PrIntB()
	//group.Wait()
	//<-end
	fmt.Println(CheckCheat(50,80))

	var c1 = make(chan bool,1)
	var c2  = make(chan bool)
	var c3  = make(chan bool)


	defer func() {
		close(c1)
		close(c2)
		close(c3)
	}()

	c1<-true

	//var group sync.WaitGroup
	//group.Add(10000)
	for range make([]struct{},10){
		go Three(c1,c2,"a")
		go Three(c2,c3,"b")
		go Three(c3,c1,"c")
	}

	 time.Sleep(10*time.Second)
	//group.Wait()

}

func CheckCheat(totalTime float64,count int )bool  {
	switch  {
	case totalTime>=0&&totalTime<20:
		return float64(count)/totalTime*10<80
	case totalTime>=20&&totalTime<30:
		return float64(count)/totalTime*20<150
	case totalTime>=30&&totalTime<60:
		return float64(count)/totalTime*30<220
	case totalTime>=60&&totalTime<120:
		return float64(count)/totalTime*60<350
	case totalTime>=120&&totalTime<300:
		return float64(count)/totalTime*120<660
	case totalTime>=300&&totalTime<600:
		return float64(count)/totalTime*300<1600
	case totalTime>=600:
		return float64(count)/totalTime*600<3000
	}
	return false
}
