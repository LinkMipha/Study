package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	UserId int  `gorm:"primary_key"`
	Phone string
	WxopenId string
	Tcreate time.Time
	Tprocess time.Time
	Balance int
	Src string
	Level int

}
type  Test struct {
	Id int
	Name string
}

var ch1 chan int
var ch2 chan int
var cha = []chan int{ch1,ch2}
var numbers = []int{1,2,3,4,5}

//func main()  {
//	select {
//	case getChan(0)<-getNumber(1):
//		fmt.Println("1th case is selected.")
//	case getChan(1)<-getNumber(2):
//		fmt.Println("2th case is selected.")
//		default:
//		fmt.Println("default!")
//
//
//	}
//
//
//}
//
//func getNumber(i int) int {
//	fmt.Printf("numbers[%d]\n", i)
//
//	return numbers[i]
//}
//func getChan(i int) chan int {
//	fmt.Printf("chs[%d]\n", i)
//	return cha[i]
//}
