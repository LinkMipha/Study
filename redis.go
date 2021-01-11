package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	fmt.Println("redis conn success")
	defer c.Close()

	result,err:=c.Do("ZREVRANGE","bicycle",0,-1)
	if err!=nil{
		fmt.Println("zrevrange错误")
	}
	switch res:=result.(type){
	case []interface{}:


	}

}

