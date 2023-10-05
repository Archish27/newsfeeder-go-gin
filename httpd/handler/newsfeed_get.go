package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type NewsfeedPostRequest struct {
	Title string `json: "title"`
	Post  string `json: "post"`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := NewsfeedPostRequest{}
		c.Bind(requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		feed.Add(item)
		c.JSON(http.StatusOK, map[string]string{
			"msg": "Created OK ",
		})
	}
}

func NewsFeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
