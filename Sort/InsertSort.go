package main

import (
	"context"
	"fmt"
)

//func InsertSort(array []int,length int){
//	for i:=1;i<length;i++{
//		for j:=i;j>0&&array[j-1]>array[j];j--{
//				array[j-1],array[j] = array[j],array[j-1]
//		}
//	}
//}
//
//func main()  {
//	array:=[]int{7,6,5,4,3,2,1}
//	length:=len(array)
//	InsertSort(array,length)
//	fmt.Println(array)
//
//
//
//	testinter:=make([]interface{},0,100)
//	testinter = append(testinter,"one","two")
//	fmt.Println(testinter)
//	for _,v:=range testinter{
//		fmt.Println(v)
//	}
//
//
//}

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main () {

	select {
	case getChan(0) <- getNumber(2):

		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):

		fmt.Println("2th case is selected.")
	default:

		fmt.Println("default!.")
	}

	context.Context.Value()
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)

	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)

	return chs[i]
}