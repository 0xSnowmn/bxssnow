package routes

import (
	"bxssnow/core"
	"fmt"
	"net/http"
	"strings"

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
	Time          string
}

func Callback(c *gin.Context) {
	var data Data

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, err)
	}
	c.JSON(http.StatusOK, data)
	//fmt.Println()
	domain := strings.SplitAfter(data.Origin, "//")
	fmt.Println(domain[1])
	core.DecodeImage(data.ScreenEncoded, domain[1])

}
