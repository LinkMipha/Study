package main

import (
	"fmt"
	"reflect"
	"time"
)

func work(s string)  {
	for i:=0;i<5;i++{
		time.Sleep(1)
		fmt.Println(s)
	}
}

func routine()  {
	ch:=make(chan int)
	go func() {
		for i:=0;i<3;i++{
			ch<-i
		}
		time.Sleep(time.Second)
	}()
	for data:=range ch{
		fmt.Println(data)
		if data==0{
			break
		}
	}

	type MyInt int
	var x MyInt = 7
	fmt.Println(reflect.ValueOf(&x))//值
	fmt.Println(reflect.TypeOf(x))

	//通过值的方式进行反射修改
	re:=reflect.ValueOf(&x)
	se:=re.Elem()
	fmt.Println(se.CanSet())
	se.SetInt(99)
	fmt.Println(x)

	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(33)
	s.Field(1).SetString("tset")

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}


}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}