package Reverse

import (
	Reverse "Test"
	"fmt"
	"testing"
)

func TestReverse(t*testing.T)  {
	head:=new(Reverse.ListNode)
	cur := head
	for i:=0;i<5;i++{
		tmp:=new(Reverse.ListNode)
		tmp.Val = i
		cur.Next = tmp
		cur = cur.Next
	}
	fmt.Println(head)
}
