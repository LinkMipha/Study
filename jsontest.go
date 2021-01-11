package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}


func js()  {

	info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}
	filePtr,err:=os.Create("test.json")
	if err!=nil{
		fmt.Println("os Create failed")
		return
	}

	defer filePtr.Close()

	encode:=json.NewEncoder(filePtr)
	err = encode.Encode(info)


	if err!=nil{
		fmt.Println("encode error")
	}else{
		fmt.Println("encode complete")
	}



}