package main

import (
	"bxssnow/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.POST("/post", routes.Callback)
	router.Run()
}
