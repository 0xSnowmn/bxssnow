package main

import (
	"bxssnow/core"
	"bxssnow/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		core.LogErrorDiscord(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.POST("/post", routes.Callback)
	router.Run(":8081")
}
