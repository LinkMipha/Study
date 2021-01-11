package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main()  {

	ch:=make([]string,0)
	one:=[]string{"one","two","three"}
	ch = append(ch,"one")
	ch = append(ch,one...)
	fmt.Println(ch)

	sh:=strings.Join(ch,",")//设置间隔符号
	fmt.Println(sh)
	hello:="hello"
	world:="world"

	//byte.buffer 是线程安全的
	var buf bytes.Buffer//buffer 方式

	//stings.builder是非线程安全的

	for i:=0;i<100;i++{
		buf.WriteString(hello)
		buf.WriteString(",")
		buf.WriteString(world)
	}
	buf.String()

	inst:=[]int{1,2,3,4}
	inst[0],inst[1] = inst[1],inst[0]
	fmt.Println(inst)
	sort.Ints(inst)
	fmt.Println(inst)
	fmt.Println("..........")

	//字符串包含中文，可以用range直接输出
	name:="中文测试"
	//butname:=[]byte(name)

	runename:=[]rune(name)

	for _,v:=range runename{
		fmt.Println(string(v))
	}

	//字符串包含中文使用rune转为utf-8 可使用普通循环访问
	for i:=0;i<len(runename);i++{
		fmt.Println(string(runename[i]))
	}

	//使用string.builder
	strbuild:=strings.Builder{}
	strbuild.WriteString("one")
	strbuild.WriteString("two")

	strbuffer:=bytes.Buffer{}
	strbuffer.WriteString("three")
	strbuffer.WriteString("four")

	fmt.Println(strbuild.String(),strbuffer.String())
}
