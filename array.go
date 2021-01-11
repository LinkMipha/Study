package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

const max = 3

type Student struct {
	Name string `json:"name"`
	Age int
}

func main()  {
	number:=[max]int{1,2,3}
	var ptrs [max]*int

	for i:=range number{
		ptrs[i] = &number[i]
	}

	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}

	str:=Student{Name: "struct name",Age: 2}
	ret:=new(Student)
	ret.Age = 1
	ret.Name = "name"
	jsret,err:=json.Marshal(ret)
	if err!=nil{
		fmt.Println("marshal出错")

	}
	jsres,err:=json.Marshal(str)
	if err!=nil{
		fmt.Println("marshal出错")

	}
	fmt.Println(jsret)
	fmt.Println(string(jsret))
	fmt.Println(string(jsres))


	//byte json数据转换为实体数据
	data:="{\"name\":\"struct name\",\"Age\":2}"
	stjson := []byte (data)
	stu:=new(Student)
	errs:=json.Unmarshal(stjson,stu)
	if errs!=nil{
		fmt.Println("逆解析失败")
	}
	fmt.Println(*stu)


	sum:=[7]int{1,2,3,4,5,6}
	fmt.Println("sum of sizeof",unsafe.Sizeof(sum))
	const name  = iota



	res:=make([]int,0)
	Add(res)
	fmt.Println(res)

}

func Add(res []int)  {
	res = append(res,1)
}

