package main

import (
	"errors"
	//"fmt"
	//"reflect"
)

/*
type error interface {
	Error() string
}

 */
func sqrt(f float64) (float64,error) {
	if f>1{
		return f,errors.New("testerror")
	}
	return 0, nil
}


/*
func main()  {
	_,st:=sqrt(3)
	if st!=nil{
		fmt.Println("error")
	}

	s := make([]int,3,5)
	//s = append(s,1)
	//s = append(s,2)
	fmt.Println(s[1:])
	typ := reflect.TypeOf(st)
	fmt.Println(typ.Name(),typ.Kind())
	var inpu string
	fmt.Scanln(&inpu)
	fmt.Println(inpu)
}

 */