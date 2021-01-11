package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
	"time"
)

func ContextTask(ctx context.Context)  {
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("接收父contxt，关闭")
			return
		default:
			go func(ctx context.Context) {
				fmt.Println("内部chnnel")
			}(ctx)
			fmt.Println("成功接收",ctx.Value("hello"))
			time.Sleep(time.Millisecond*1000)
		}
	}
}

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:

		}
	}
}
func main()  {

	f, err := os.OpenFile("/tmp/cpu2.pprof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 8; i++ {
		go logicCode()
	}
	ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*1000*5)

	//deadline 以现在的时间为界限
	//ctx,cancel:=context.WithDeadline(context.Background(),time.Now().Add(time.Millisecond*10))


	ctxValue:=context.WithValue(ctx,"hello","world")
	go ContextTask(ctxValue)
	go ContextTask(ctxValue)

	time.Sleep(time.Millisecond*10000)
	cancel()
	fmt.Println(ctx.Err())
	http.ListenAndServe("localhost:6060", nil)
	time.Sleep(time.Second*10)
	os.Exit(0)

}
