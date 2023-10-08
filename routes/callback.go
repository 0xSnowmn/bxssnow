package routes

import (
	"bxssnow/core"
	"net/http"

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
	core.DecodeImage(data.ScreenEncoded)

}
