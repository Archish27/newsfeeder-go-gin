package main

import (
	"fmt"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)


func main()  {
	
	fmt.Println("Newfeeder Application")

	r := gin.Default()

	feed := newsfeed.New()
	
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	
	r.Run()
}
