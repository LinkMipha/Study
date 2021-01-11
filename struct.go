package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
	Id int
}

func (a A)isEmpty()bool  {
	return reflect.DeepEqual(a,A{})
}

func main()  {
	var st A
	fmt.Println(st.isEmpty())
	st.Id = 11
	fmt.Println(st.isEmpty())

	ret:=[]int{1,2,3}
	ret = ret[:len(ret)-1]
	fmt.Println(ret)
}