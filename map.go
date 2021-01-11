package main

import (
	"fmt"
	"reflect"
)

//import "fmt"

func ma()  {

	st:=map[int]string{1:"one",2:"two"}

	var ma=map[int]string{1:"one",2:"two",3:"three",4:"four",
		5:"five",
	}
	delete(ma,2)
	delete(ma,5)
	for _,j:=range ma{
		fmt.Println(j)
	}

	if v,help:=ma[3];help{
		fmt.Println(v,help)
	}
	if reflect.DeepEqual(ma,st){
		fmt.Println("st==ma")
	}else {
		fmt.Println("st!=ma")
	}
	test:=new(testa)
	test.Name = 1
	test.add(2)
	fmt.Println(test.Name)

	var se interface{}
	se = 3
	pr(se)
	pr("test")
	pr(3)

	v,ok:=se.(int)
	if ok{
		fmt.Println(v)
		fmt.Println(ok)
	}

	switch se.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("其他")
	}

	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))

	f := []int{1, 2, 3}
	fmt.Println("cap of f before:", cap(f))
	f = append(f, 4)
	fmt.Println("cap of f after:", cap(f))

}

func pr(a interface{})  {
	fmt.Println(a)
}


type testa struct {
	Name int
}

func (e *testa)add(b int)  {
	e.Name+=b
}