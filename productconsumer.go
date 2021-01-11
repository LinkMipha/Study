package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

var done chan int = make(chan int)
var bufchan chan int = make(chan int,100)

func  Write(max int)  {
	for i:=0;i<max;i++{
		bufchan<-i
		time.Sleep(time.Millisecond*100)
	}

}

func Clo()  {
	close(done)
}

func Read(max int)  {
	for{
		st:=<-bufchan
		fmt.Println("chan read",st)
	}
	once.Do(Clo)
}

func TestreadAndwrite(max int)  {
	go Write(max)
	go Read(max)

	value,ok:=<-done
	if ok{
		fmt.Println(value)
	}else{
		fmt.Println("程序结束")
	}
	close(bufchan)

}





func main()  {
	TestreadAndwrite(100)
	
}