package main

import (
	"crypto/md5"
	"fmt"
)

//import "fmt"

func sum(n int) int {
	if n<2 {
		return n
	}
	return sum(n-1)+sum(n-2)
}
/*
func main()  {
	st:=sum(42)
	fmt.Println(st)
}

 */

func main()  {
	st:=testfunc(0)
	//fmt.Println(st)

	fmt.Printf("%T",st)
	f := []string{}
	for i := 0; i < 16; i++ {
		f = append(f, "1.1", "2.2", "3.3", "4.4", "5.5")
		fmt.Print(cap(f), " ")
	}
	str :="123456"
	h := md5.New()
	h.Write([]byte(str))
	fmt.Println(h.Sum(nil))
	fmt.Println(md5.New().Sum(nil))
	fmt.Println(md5.New().Sum(nil))

}

func testfunc(st int) int {
	fmt.Println("test func ")
	return st
}
