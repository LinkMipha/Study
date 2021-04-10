package main

import (
	"flag"
	"fmt"
)

func main()  {
	var user string
	var password string

	flag.StringVar(&user,"u","","默认值")
	flag.StringVar(&password,"pwd","","密码")
	port:=flag.Int("port",3306,"端口 返回地址格式，打印需要解引用")
	st:=flag.Bool("b",true,"备注")


	flag.Parse()//访问之前需要这一步，从arguments 中解析注册到 flag，不加只是地址关联并没有
	fmt.Printf("user:%v password:%v,port:%v,bool:%v",user,password,*port,*st)
}