package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	//UTC 世界标准时间
	//GMT 格林尼治平时
	//CST CST可以同时表示美国，澳大利亚，中国，古巴四个国家的标准时间。
	//北京时间所属时区: UTC/GMT +8

	//字符串转换比较慢
	beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-09-27 19:00:00", time.Local)
	fmt.Println(beginTime)
	beginTime = time.Date(2020,9,27,19,0,0,0,time.Local)
	fmt.Println(beginTime)
	//UTC 时间可能会有差值
	startTime,_:=time.Parse("2006-01-02 15:04:05","2020-09-27 19:00:00")
	fmt.Println(startTime)

	timenow:=time.Now()
	fmt.Println("timeNow",timenow)
	//获得当前的时间
	//fmt.Println(timenow.Uninx().Format("2006-01-02 15:04:05"))
	fmt.Println(timenow.Unix())
	fmt.Println(beginTime.Add(time.Second*10).Equal(beginTime))

	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n",time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)
	//preTime,_:=time.ParseDuration("-24h")
	fmt.Println("ststst",time.Now().Add(-time.Hour*time.Duration(24)))
	beginTime = time.Now()
	time.Sleep(600*time.Millisecond)
	endTime := time.Now()
	//整除截断保留整数
	fmt.Println("time sub",endTime.UnixNano()/1e9,(endTime.Sub(beginTime)).Nanoseconds()/1e6)


	// 获取50秒前的时间，方式1
	st,_ := time.ParseDuration("-50s")
	fmt.Println("50秒前的时间：",time.Now().Add(st))

	// 获取1分钟前的时间，n秒前则是time.Second * -n，方式2
	t := time.Now().Add(time.Minute * -1)
	fmt.Println("一分钟前的时间：",t)

	//获取1小时前的时间
	sth,_ := time.ParseDuration("-1h")
	fmt.Println("1小时前的时间：",time.Now().Add(sth))

	// 获取2天前的时间
	oldTime := time.Now().AddDate(0, 0, -2)
	fmt.Println("time add second",oldTime.Add(time.Second*10))
	//获取两个月前的时间gip
	//oldTime := time.Now().AddDate(0, -2, 0)

	fmt.Println("set"+strconv.FormatInt(124, 10))

	timen:=time.Now()
	time.Sleep(1*time.Second)
	elapsed:=time.Since(timen)
	fmt.Println("times ",elapsed)



}