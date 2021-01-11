package main

type T struct {
	ls []int
	v  int
}

func foo(t T)  {
	t.ls[0] = 999
	t.v = 666
}


//func sync() {
//	str :=string("[1,2,3,4,5]")
//	keys := make([]int64,0)
//	json.Unmarshal([]byte(str), &keys)
//
//	fmt.Println(keys)
//	slice:=[]int{10,20,30,40,50}
//	newslice:=slice[1:2:2]//容量为1 重新创建底层数组
//	newslice = append(newslice,60)
//	var err error
//	err = errors.New("error测试")
//	fmt.Println(err)
//	fmt.Printf("%#v\n",slice)//显示格式  p地址  T类型
//
//	fun:= func(str string)string {// b f  左右
//		fmt.Printf("%v\n",str)//p n 上下
//		return str
//	}
//	fun("test")
//
//	t :=T{ls: []int{1,2,3}}
//	foo(t)
//	fmt.Println(t)
///*
//	value,ok:=interface{}(slice).([]int)
//	if !ok{
//		fmt.Println("errors")
//	}
//	fmt.Println(value)
// */
//	s := fmt.Sprintf("是字符串 %s ","string")
//	fmt.Println(s)
//
//	for i:=0;i<100;i++ {
//		if i>10{
//			continue
//		}
//		fmt.Println("循环中")
//	}
//
//	strslice :=strings.Split("hello world test apple"," ")
//	fmt.Println(strslice)
//
//	for i,v:=range strslice{
//		fmt.Printf("%d : %s\n",i,v)
//	}
//
//	type animal struct {
//		Result []struct{
//			Name string
//			Order string
//		}
//	}
//
//	var jsstr = []byte(	`{"Result":[
//	{"Name": "Platypus", "Order": "Monotremata"},
//	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
//]}`)
//
//	jsres:=animal{}
//	jserr:=json.Unmarshal(jsstr,&jsres)
//	if jserr!=nil{
//		fmt.Println("error",jserr)
//	}
//	fmt.Printf("%+v\n",jsres)
//
//	testsort:=[]int{5,2,1,4,3}
//	sort.Ints(testsort)
//	fmt.Println(testsort)
//
//	/*
//	data:=[]string{"one","two","three"}
//	for _,v:=range data{
//		st:=v
//		go func() {
//			fmt.Println(st)
//		}()
//	}
//	 */
//	type Stu struct {
//		name string
//	}
//
//	ma:=map[string]Stu{
//		"x":{"one"},
//	}
//	fmt.Println(ma)
//	//ma["x"].name = "test"  找不到map["x"]对应的struct值
//
//	sst:=ma["x"]
//	sst.name = "two"
//	ma["x"]=sst
//
//	fmt.Println(ma)
//	slicemap:=[]Stu{{"one"}}
//	slicemap[0].name  = "change"
//
//	mapoint:=map[string]*Stu{
//		"x":{"point"},
//	}
//	mapoint["x"].name = "change point"
//	fmt.Println(&mapoint)
//	fmt.Printf("%+v\n",mapoint)
//
//	resslice := []int{1, 2, 3, 4}
//	for _, v := range resslice {
//		fmt.Println(v)
//		defer func() {
//			fmt.Println(v)
//		}()
//	}
//
//}
//func fun1(value int)  {
//	fmt.Println(value)
//}

