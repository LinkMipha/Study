package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func iou()  {
	url:="http://www.baidu.com"
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
