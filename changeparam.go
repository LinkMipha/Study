package main

import (
	"errors"
	"fmt"
)


//可变参数在尾部
func changeParam(str string,a ...string)  {
	fmt.Println(str)

	for _,v:=range a{
		fmt.Println("change param",v)
	}
}

func main()  {
	stone:=make([]string,10)
	st:=[]string {"two","three","four","five"}
	stone = append(stone,st...)

	//可变参数切片可以传入
	changeParam("one",stone...)
	two:=make([]string,len(st))
	copy(two,st)
	changeParam("two",two...)
	fmt.Println(correctRelease())
	//fmt.Println(f2())
	fmt.Println(*f3())
}


func correctRelease() (err error) {
	defer func() {
		err = errors.New("error")
	}()
	return nil
}

//func f1() (i int) {
//
//	i = 1
//	defer func(){
//		i++
//	}
//	return i
//}

func f2() int {
	i := 1
	defer func(){
		i++}()
	return i

}


func f3() *int {
	i := 1
	defer func(){i++}()

	return &i

}

func test() (res int) {
	tmp := 1
	defer func() {
		tmp++
	}()
	return tmp
}
