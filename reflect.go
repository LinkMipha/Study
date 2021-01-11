package main

import (
	"fmt"
	"reflect"
)

type Reflect struct {
	Name string `json:"name" test:"tag"`
	PassWord string
	Three int

}

func (m Reflect)Hello(v int)  {
	fmt.Println("hello",v)
}

func (m Reflect)World()  {
	fmt.Printf("world")
}

func main()  {
	//st:=reflect.DeepEqual()
	s:=Reflect{"name","password",0}
	//类型
	ty:=reflect.TypeOf(s)
	//值
	va:=reflect.ValueOf(s)

	//kind判断类型
	if ty.Kind()!=reflect.Struct{
		fmt.Println("error not struct")
	}
	//type 相关的操作
	//结构体名字
	fmt.Println(ty.Name())
	//字段数
	fmt.Println(ty.NumField())

	for i:=0;i<ty.NumField();i++{

		//获取字段具体的值
		value:=va.Field(i).Interface()

		//获取单个字段
		filed:=ty.Field(i)

		//filed.Tag可以获取tag，通过tag进行参数的校验
		fmt.Printf("第%d 个字段是%s,%s ,值为%v \n",i,filed.Name,filed.Type,value)

		//打印后面的tag标签
		fmt.Println(filed.Tag)
	}

	//获取方法相关
	fmt.Println(ty.NumMethod())
	for i:=0;i<ty.NumMethod();i++{
		//方法名字和类型
		method:=ty.Method(i)
		fmt.Printf("方法名:%s 类型:%v\n",method.Name,method.Type)
	}



	//通过反射修改内容,必须取出地址
	add:=&Reflect{"name","password",0}


	v:=reflect.ValueOf(add)
	// 修改值必须是指针类型否则不可行
	if v.Kind() != reflect.Ptr {
		fmt.Println("不是指针类型，没法进行修改操作")
		return
	}

	//获取指针
	ptr:=v.Elem()

	//获取具体字段的封装
	name:=ptr.FieldByName("Name")
	if name.Kind()!=reflect.String{
		fmt.Printf("类型出错 \n")
		return
	}

	//修改具体的值
	name.SetString("change")
	fmt.Println(add.Name)

	//简单的修改整形
	ret:=999
	addr:=reflect.ValueOf(&ret)
	addr.Elem().SetInt(123)
	fmt.Println(ret)




	//通过反射调用方法

	//有参方法
	v1:=reflect.ValueOf(s)

	method1:=v1.MethodByName("Hello")
	arg:=[]reflect.Value{reflect.ValueOf(1)}
	method1.Call(arg)


	//调用无参数方法
	method2:=v1.MethodByName("World")
	var arg2 []reflect.Value
	method2.Call(arg2)

}