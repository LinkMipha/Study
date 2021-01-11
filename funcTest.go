package main

import (
	"crypto/rand"
	"fmt"
	"time"

	//"git.in.codoon.com/third/go-redis"
)
type RedisConfig struct {
	RedisConn      string
	RedisPasswd    string
	ReadTimeout    int
	ConnectTimeout int
	WriteTimeout   int
	IdleTimeout    int
	MinIdle        int
	MaxIdle        int
	MaxActive      int
	RedisDb        string
	MaxRetries     int
}
/*
type RedisClient struct {
	*redis.Client
}
var conf RedisConfig

func initRedis(conf *RedisConfig)()  {

}*/



func printNumber(from, to int, c chan int) {
	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep(1 * time.Millisecond)
	}
	c <- 0
}


func main(){
	b := make([]byte, 16)
	rand.Read(b)
	//编码
	//res:=base64.StdEncoding.EncodeToString(b)

	st:=[]int{1,2,3}
	ret:=make([]int,0)
	ret =append(ret,st...)
	ret = ret[:len(ret)]
	fmt.Println(ret)

	c := make(chan int, 3)
	go printNumber(1, 3, c)
	go printNumber(4, 6, c)
	_ = <- c
	_ = <- c
	_ = <- c

}

func mergess(nums1 []int, m int, nums2 []int, n int) []int {
	left:=0
	right:=0
	res:=make([]int,0)
	for left<m&&right<n{
		if nums1[left]<nums2[right]{
			res = append(res,nums1[left])
			left++
		}else{
			res = append(res,nums2[right])
			right++
		}
	}
	if left>=m{
		for right<n{
			res = append(res,nums2[right])
			right++
		}
	}else{
		for left<m{
			res = append(res,nums1[left])
			left++
		}
	}
	nums1 =  res
	return nums1
}
