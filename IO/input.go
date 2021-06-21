package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//bufio读取
	reader:=bufio.NewReader(os.Stdin)
	//直到遇到匹配的字符
	str,_:=reader.ReadString('\n')
	fmt.Println("input",str)

	//直接方式读取
	var scanStr string
	fmt.Scanln(&scanStr)
	fmt.Printf("scanln string : %s\n",scanStr)
}