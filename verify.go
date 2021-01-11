package main

import "fmt"

func verify(array []int,left,right int)bool  {
	if right<=left{
		return true
	}
	i:=left
	for i<right&&array[i]<array[right]{
		i++
	}
	j:=i
	for j<right{
		if array[j]<array[right]{
			return false
		}
		j++
	}
	le:=true
	ri:=true
	if i>left{
		le = verify(array,left,i-1)
	}

	if i<right{
		ri = verify(array,i,right-1)
	}
	return le&&ri
}

func main()  {
	array:=[]int{4, 8, 6, 12, 16, 14, 10}
	fmt.Println(verify(array,0,len(array)-1))
}