package main

import "fmt"

type Str struct {
	str string
}

func main()  {
	//str:=new(Str)
	fmt.Println(&Str{
		str: "tset",
	})

}

//func Test(str string,ret int)*Str  {
//	 str:=new(Str)
//	return &str{}
//}