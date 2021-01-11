package main

import (
	"fmt"
	"time"
)

var notify  = make(chan int)
var buf = make(chan int,100)

func product()  {
	<-notify
	fmt.Println("接收到生产信号")
	for i:=0;i<100;i++{
		buf<-i
		fmt.Println("生产产品为",i)
	}
	fmt.Println("生产结束")
	close(notify)//关闭后缓冲仍有数据可以读取
}

//一个消费者
func consumer()  {
	time.Sleep(time.Millisecond*100)
	fmt.Println("发送开始信号")
	notify<-1
	time.Sleep(1 * time.Second)
	for{
		if co,ok:=<-buf;ok{
			fmt.Println("接收",co)
		}else{
			fmt.Println("接收不到")
			break
		}
	}
}

func consumermore()  {
	time.Sleep(time.Millisecond*100)
	fmt.Println("发送开始信号")
	notify<-1
	time.Sleep(1 * time.Second)
	for i:=0;i<5;i++{
		//如果不加参数，形成闭包，输出最后一位
		go func(i int) {
			for{
				//time.Sleep(time.Millisecond*100)
				if value,ok:=<-buf;ok{
					fmt.Printf("第%d个消费者，消费为%d\n",i,value)
				}else{
					fmt.Printf("结束")
					break
				}
			}
		}(i)
	}

}

func main()  {

	go product()
	go consumermore()
	time.Sleep(5 * time.Second)
}