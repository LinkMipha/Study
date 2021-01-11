package main

import  "fmt"

func main()  {
	var ret interface{}
	ret = 12
	value,ok:=ret.(string)
	if !ok{
		fmt.Println(value)
	}else{
		fmt.Println(ok)
	}

	//配合switch使用断言
	switch t:=ret.(type){
	case string:
		fmt.Println("string",t)
	case int64:
		fmt.Println("int64",t)
	default:
		fmt.Println("other",t)
	}


}