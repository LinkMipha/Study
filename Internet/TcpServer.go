package main

import (
	"fmt"
	"net"
)

func main()  {
	listen,err:=net.Listen("tcp",":8080")
	if err!=nil{
		fmt.Printf("listen error :%v\n",err)
		return
	}
	for{
		conn,accErr:=listen.Accept()
		if accErr!=nil{
			fmt.Printf("accept error: %v\n",accErr)
			break
		}
		go HandleConn(conn)
	}
}


func HandleConn(conn net.Conn) {
	var read []byte
	//read:=make([]byte,256)
	for  {
		read=make([]byte,256)
		length,err:=conn.Read(read)
		if err!=nil{
			fmt.Println("Read error")
			break
		}
		fmt.Println(string(read),"read message length",length)
		length,err=conn.Write(read)
		fmt.Println("write message",string(read))
	}
}

