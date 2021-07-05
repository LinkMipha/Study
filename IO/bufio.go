package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//scanner读取缓冲区数据
var scanner *bufio.Scanner

func main() {
	//input := "foo  bar   baz"
	//scanner := bufio.NewScanner(strings.NewReader(input))
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

	scanner = bufio.NewScanner(os.Stdin)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))//定义缓冲区大小

	scanner.Scan() //扫描数据
	fmt.Println(scanner.Text())//打印一行数据
}


//读取一行数据
func readLine() []int {
	scanner.Scan()
	line := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	res := make([]int, len(line))
	for i, num := range line {
		x, _ := strconv.Atoi(num)
		res[i] = x
	}
	//return res
   //i:=-1e9
   return []int{}
}

