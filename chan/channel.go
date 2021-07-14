package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
channel 传入值没有关闭就使用range调用会死锁


 */
func tet()  {

	var ch chan int
	//func1
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	//func2
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}

}


func main()  {
	/*ont:=sync.WaitGroup{}
	ont.Add(10)
	st:=make([]int,11)
	for i:=0;i<10;i++{
		go func() {
			st[i] = i
		}()
		ont.Done()
	}
	ont.Wait()*/


	ch := make(chan int)
	close(ch)
	//ch <- "good" //向关闭的channel中写会panic

	//向关闭的channel中读取，会读取默认值
	i := <- ch
	fmt.Println("close channel",i)

	//判断是否关闭
	value,ok:=<-ch
	if ok{
		fmt.Println(value)
	}else{
		fmt.Println("channel close")
	}

//	关闭未初始化和已经关闭的chan 会panic   读取未初始化的chan 阻塞死锁 读取已经关闭的chan 默认值
	/*var a chan int
	 //a:=make(chan int)
	go func() {
		<-a
	}()
	close(a)
	fmt.Println(<-a)*/



}