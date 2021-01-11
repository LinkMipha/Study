package main

type Result struct {
	Id int
	name string
	count int
}
func biranySearch(array []int,target int)int{
	var low int = 0
	var high int = len(array)
	for low<=high{
		mid:=low+(high-low)/2
		if array[mid]==target{
			return mid
		}else if array[mid]>target {
			high = mid - 1
		}else if array[mid]<target{
			low = mid+1
		}
	}
	return -1
}

//func main() {
//	array:=[]int{2,4,6,8,10}
//	fmt.Println(biranySearch(array,6))
//	for loop := true; loop; loop = false{
//		fmt.Println(loop)
//	}
//
//	st:=make([]Result,0)
//	st = append(st, Result{
//		Id: 1,
//		name: "one",
//		count: 1,
//	})
//
//	st = append(st, Result{
//		Id: 2,
//		name: "two",
//		count: 2,
//	})
//
//	st = append(st, Result{
//		Id: 3,
//		name: "three",
//		count: 3,
//	})
//
//	for _,v:=range st{
//		if v.Id==2{
//			v.name = "change"
//		}
//	}
//	fmt.Println(st[1])
//	var flag bool
//	fmt.Println(flag)
//
//	ones:=[]string{"1","2","3","4"}
//	ones = ones[2:len(ones)-1]
//	stringtest:=strings.Join(ones,"+")
//	fmt.Println(stringtest)
//	str:="hello world"+"测试"
//	fmt.Println(len([]rune(str)))
//}