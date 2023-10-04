package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	router := gin.Default()
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.POST("/post", func(c *gin.Context) {
		var data Data
		c.BindJSON(&data)
		c.JSON(http.StatusOK, data)
		fmt.Println(data)
	})
	router.Run()
}
