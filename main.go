package main

import (
	"bxssnow/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.StaticFile("/mm.html", "./injector/mm.html")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	router.POST("/post", routes.Callback)
	router.Run(":8081")
}
