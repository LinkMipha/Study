package Test

import (
	"fmt"
	"testing"
)

func TestHello(t*testing.T)  {
	a,b:=1,2
	c:=Hello(a,b)
	fmt.Println(c==1)
}

func TestWrold(t*testing.T)  {
	a,b:=5,1
	c:=Wrold(a,b)
	fmt.Println(c==4)

}



