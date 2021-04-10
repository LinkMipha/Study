package main

import (
	"flag"
	"fmt"
)

func main()  {
	var user string
	var password string

	flag.StringVar(&user,"u","root","备注")
	flag.StringVar(&password,"pwd","LinkMipha","密码")
	port:=flag.Int("port",3306,"端口 返回地址格式，打印需要解引用")
	st:=flag.Bool("b",true,"备注")
	fmt.Printf("user:%v password:%v,port:%v,bool:%v",user,password,*port,*st)
}