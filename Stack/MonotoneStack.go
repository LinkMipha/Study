package main

import "fmt"

//单调栈

func main()  {
	array:=[]int{1,2,3,4}
	test(array)
	//传入切片
	fmt.Println(array)
}

func test(arr []int)  {
	arr[0]=100
}