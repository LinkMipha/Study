package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	require:=make(chan int,5)

	for i:=0;i<1;i++{
		require<-i
	}

	close(require)

	//设置定时器控制读取速率
	ti:=time.Tick(time.Duration(1)*time.Second)

	//直接range只会读取值，缓冲区无数据会读取默认零值，chan关闭，则退出range
	for i:=range require{
		<-ti
		fmt.Printf("i: %d time: %v\n",i,time.Now())
	}

	//DeepEqual()
	fmt.Println(reflect.ValueOf(ti))
	fmt.Println(reflect.TypeOf(ti))

	hashmap:=make(map[int]int,0)
	hashmap[0] = 99
	if value,ok:=hashmap[0];ok{
		fmt.Println(value)
	}
	st:=int(^uint(0)>>1)
	fmt.Println(st,^st)


}