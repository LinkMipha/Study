package main

import (
	"fmt"
	"time"
)



func main()  {

	var outtime  = time.Millisecond*10
	var capacity = 100
	var tockenBucket  = make(chan struct{},capacity)
	bucket:= func() {
		ticker:=time.NewTicker(outtime)
		for{
			select {
				case <-ticker.C:
					select {
					case tockenBucket<- struct{}{}:
					default:
					}
					fmt.Printf("time:%v  capacity:%d\n",time.Now(),len(tockenBucket))
					fmt.Printf("hello world\n")
			}
		}
	}

	go bucket()
	time.Sleep(time.Hour*1)
}
