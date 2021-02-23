package main

import (
	"encoding/json"
	"fmt"
)

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
	Struct json.RawMessage `json:"stu"` //保留json的原值
	//如果用interface接收，通常结构体被转换为map[string]interface{}类型

}

//用原始数据再次进行unmarshal可以解析出结果
type Stu struct {
	One interface{}
	Two interface{}
}

func main()  {

	stu:=StuRead{"link",12,122,1,"class","test",{\"naME\":\"1班\",\"GradE\":3}}

	marshal,err:=json.Marshal(stu)
	if err!=nil{
		fmt.Println(err.Error())
	}
	//序列化后转换为字符串方便查看
	fmt.Println(marshal,"string后",string(marshal))

	//构造字符串，直接解析,字符串用大括号，另外引号需要进行转义
	str:="{\"name\":\"link\"}"
	var stu2 StuRead
	//unmarshal 把[]byte转换为指针数据
	err =json.Unmarshal([]byte(str),&stu2)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(stu2)


}
