package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	//rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(10))

	fmt.Println(rand.Intn(10))

	for i:=0;i<1000;i++{
		fmt.Println(rand.Intn(500)+100)
	}
}