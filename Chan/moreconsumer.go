package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//多生产单个消费者
/*var dataCh = make(chan int,1)
var doneCh = make(chan int)
func pro()  {
	for i:=0;i<1000;i++{
		go func(i int) {
			for {
				select {
				case <-doneCh:
					fmt.Println("接收到停止发送的信号")
				case dataCh<-i:
					fmt.Printf("第%d结束发送\n",i)
				}
			}
		}(i)
	}
}

func consu()  {
	for{
		if value,ok:=<-dataCh;ok{
			fmt.Println("输出",value)
			if value==998{
				close(doneCh)
				return
			}
		}else{
			fmt.Println("输出结束")
			break
		}
	}
}

*/




//多生产多消费
type Product struct {
	name int
	value int
}

func Produ(wg *sync.WaitGroup,product chan<- Product,name int,stop *bool)  {
	for !*stop{
		pro:=Product{name: name,value: rand.Int()}
		product<-pro
		fmt.Printf("producer %v produce a product: %#v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}



func consume(wg *sync.WaitGroup,product <-chan Product,name int)  {
	for pro:=range  product{
		fmt.Printf("consumer %v consume a product: %#v\n", name, pro)
	}
	wg.Done()
}


func main()  {
	stop:=false
	product:=make(chan Product,10)
	wgc:=sync.WaitGroup{}
	wgp:=sync.WaitGroup{}


	for i:=0;i<5;i++{
		go Produ(&wgp,product,i,&stop)
		go consume(&wgc,product,i)
		wgc.Add(1)
		wgp.Add(1)
	}
	time.Sleep(time.Duration(1) * time.Second)
	stop = true
	wgp.Wait()
	close(product)
	wgc.Wait()

	time.Sleep(time.Second*5)
}