package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var flag = 1
var group sync.WaitGroup

var tmp = time.Tick(time.Second*time.Duration(2))

func Other(flag *int)  {
	defer group.Done()
	*flag = 0
}

func Upload(flag *int)int {
	defer func() {
		go Other(flag)
	}()
	return 0
}


func main()  {
	group.Add(1)
	Upload(&flag)
	group.Wait()
	fmt.Println(flag)

	startTime:=time.Now()
	keyDay := fmt.Sprintf("distance:day:%d-%02d-%02d", startTime.Year(), startTime.Month(), startTime.Day())
	fmt.Println(keyDay)
	//ride:distance:day:2021-01-27
	var st = 29.120
	fmt.Println( float64(   int(st/1000*100)   )/100  )

}
func formatFloat(val float64) float64   {
	//return / 100
	return float64(int(math.Floor(val*100+0.5))) / 100
}
