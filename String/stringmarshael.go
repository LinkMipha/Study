package main

import (
	"encoding/json"
	"fmt"
)

type Marshael struct {
	Name string
	Id int
	Path string
}

type Data struct {
	Test int
}

func main()  {
	msg:=Marshael{}
	data,err:=json.Marshal(msg)
	if err!=nil{
		fmt.Println(data)
	}

	unData:=Data{}
	err = json.Unmarshal(data,&unData)
	if err!=nil{
		fmt.Println(err)
	}

}

