package routes

import (
	"bufio"
	"bxssnow/core"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/structs"
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
	domain := strings.SplitAfter(data.Origin, "//")
	core.Optmize(data.ScreenEncoded, domain[1])
	core.Msg = "XSS Fired! at " + data.URL
	core.S()
}

func readTpl() string {
	str := ""
	file, err := os.Open("./tpl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f := structs.Names(&Data{})
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		for _, name := range f {
			strings.Replace(scanner.Text(), fmt.Sprintf("{{%s}}", name), Data.Dom, 20)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return str
}
