package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"git.in.codoon.com/third/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main()  {
	router := gin.Default()
	router.PUT("/ping", func(c *gin.Context) {
		req:=BackendUploadF2WatchBannerReq{}

		if !c.Bind(&req) {
			fmt.Println("bind error")
		}
		c.JSON(200, gin.H{
			"one":req.Id,
			"two": req.PicUrl,
			"three":req.Email,
			"message": "pong",
		})
	})

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})


	router.POST("/post", func(c *gin.Context) {
		message:=c.PostForm("name")
		one:=c.DefaultPostForm("age","12")

		c.JSON(200,gin.H{
			"status":"200",
			"message":message,
			"test":one,

		})
	})
	timeStamp := time.Now().Unix()
	fmt.Println(timeStamp)
	var b []byte  = nil
	fmt.Println("md5")
	fmt.Println(Md5(b))

	testurl:= url.Values{}
	testurl.Set("test","st")

	rawUrl:="http://120.26.10.118:2048/v2/equipment/get_k2_watch_style"
	reqUrl,err:=url.ParseRequestURI(rawUrl)
	if err!=nil{
		fmt.Printf("url.parserequesturl %v",err)
		return
	}

	resp,err:=http.Get(reqUrl.String())
	if err!=nil{
		fmt.Printf("url.get %v",err)
		return
	}

	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Printf("url.ReadAll %v",err)
		return
	}
	fmt.Println(string(body))


	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run(":8000")
}

func Md5(data []byte) string {
	md5:=md5.New()
	md5.Write(data)
	md5data:=md5.Sum([]byte(""))
	return hex.EncodeToString(md5data)
}


type BackendUploadF2WatchBannerReq struct {
	Id int64 `json:"id" form:"id"`
	//banner图文件key
	PicUrl string `json:"pic_url"`
	Email  string `json:"email" form:"email"`
}




