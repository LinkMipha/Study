package main


import "fmt"

func main()  {

	//切片扩容机制，容量小于1024 翻倍，大于则每次增加1/4，append会内存对齐，所以容量不严格按照此规则
	arr:=[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 :=arr[2:5]
	//cap 为8，不加第三个参数，默认为底层数组结尾，
	fmt.Println(s1,len(s1),cap(s1))


	s2 := s1[2:6:7]
	fmt.Println(s2,len(s2),cap(s2)) //len 为6-2  cap为7-2
	


}
