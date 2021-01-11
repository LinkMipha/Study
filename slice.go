package main

import "fmt"

/*
func main() {
	//var numbers = make([]int,4,5)
	var slice[] string
	slice = append(slice,"appendone")
	slice = append(slice,"appendtwo")
	slice = append(slice,"appendthree")
	slicetwo :=slice[:4]//slice 的引用
	//slicetwo:=make([]string,len(slice),(cap(slice))*2)
	//copy(slicetwo,slice)
	print(slice)
	print(slicetwo)

}

 */

func print(x []string)  {
	if x==nil {
		fmt.Println("nil slice")
	}
	fmt.Println(cap(x),len(x),x)


}
