package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

//trace 工具使用
func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}

func main() {

	runtime.GOMAXPROCS(1)

	f, _ := os.Create("trace.output")
	defer f.Close()

	//使用trace进行分析
	_ = trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(&wg, i)
	}

	wg.Wait()
	_cmpDate := time.Now().Local()
	cmpDate := &_cmpDate
	t, _ := time.ParseInLocation("2006-01-02",  cmpDate.Format("2006-01-02"), time.Local)

	weekDay := t.Weekday()
	fmt.Println(weekDay)

	stack:=make([]byte,0)

	fmt.Println(len(stack))
}

