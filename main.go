package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFile("/extractor.js", "./injector/extractor.js")
	router.Run()
}
