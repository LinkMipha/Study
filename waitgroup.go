package main

import (
	"fmt"
	"sync"
	"time"
)

func wgone(wg *sync.WaitGroup)  {
	fmt.Println("one")
	wg.Done()
}

func groupadd(wg *sync.WaitGroup)  {
	fmt.Println("增加waitgroup信号量")
	wg.Add(1)
}

func  main()  {

	var wg sync.WaitGroup
	wg.Add(2)
	go wgone(&wg)
	go wgone(&wg)
	go groupadd(&wg)
	time.Sleep(time.Second*2)
	//wg.Add(-2)

	wg.Wait()
	//wg.Add(1)
	wg.Wait()
}