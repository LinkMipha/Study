package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type IdTest struct {
	Id []string `json:"id"`
}

func main()  {

	////按照行读取
	//file,err:=os.Open("/Users/codoon/go/src/Test/Te.text")
	//if err!=nil{
	//	fmt.Printf("open file failed :%v",err)
	//	return
	//}
	//defer file.Close()
	//
	//ids:=IdTest{}
	//bu:=bufio.NewReader(file)
	//for{
	//	a,_,c := bu.ReadLine()
	//	if c==io.EOF{
	//		fmt.Printf("读取结束")
	//		break
	//	}
	//	ids.Id = append(ids.Id,string(a))
	//}
	//
	////去除开头空格
	//for i,_:=range ids.Id{
	//	ids.Id[i] = ids.Id[i][1:]
	//}
	//fmt.Println(ids.Id)



	//直接读取全文

	one,err:=ioutil.ReadFile("/Users/codoon/go/src/Test/Te.text")
	if err!=nil{
		fmt.Printf("readfile open failed:%v",err)
		return
	}
	//一个多个空格生成切片
	str:=strings.Fields(string(one))
	fmt.Println(str)


	writefile, err := os.OpenFile("TestCsv.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err!=nil{
		fmt.Println("open file failed")
		return
	}
	defer writefile.Close()
	writefile.WriteString("\xEF\xBB\xBF")// 写入UTF-8 BOM，防止中文乱码
	w := csv.NewWriter(writefile)



	length:=len(str)
	rows:=make([][]string,length)
	for i:=0;i<length;i++{
		rows[i] = []string{str[i],}
	}

	rows = append([][]string{{"跑团id"}},rows...)


   	w.WriteAll(rows)
   // 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
    w.Flush()
}