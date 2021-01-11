package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	name int
	value int
}

func Producter(wg *sync.WaitGroup,products chan <-Product,name int,flag *bool)  {
	for !*flag{
		product:=Product{name: name,value: rand.Int()}
		products <-product
		fmt.Printf("product %v product %+v\n",name,product)
		time.Sleep(time.Duration(200+rand.Intn(1000))*time.Millisecond)
	}
	wg.Done()
}

func Consumer(wg *sync.WaitGroup,products <-chan Product,name int){
	for consumer:=range products{
		fmt.Printf("consumer %v product %+v\n",name,consumer)
		time.Sleep(time.Duration(200+rand.Intn(1000))*time.Millisecond)

	}
	wg.Done()
}

func main()  {
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	products:=make(chan Product,10)
	flag:=false
	for i:=0;i<5;i++{
		go Producter(&wgp,products,i,&flag)
		go Consumer(&wgc,products,i)
		wgp.Add(1)
		wgc.Add(1)
	}
	time.Sleep(time.Duration(1)*time.Second)
	flag = true
	wgp.Wait()
	close(products)
	wgc.Wait()
}