package core

import (
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"time"

	"image/png"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/h2non/bimg"
)

func DecodeImage(b64 string, domain string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	pngFilename := domain + ".png"
	f, err := os.Create("./screenshots/" + pngFilename)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	ff, _ := os.Open("./screenshots/" + pngFilename)
	buffer, err := io.ReadAll(ff)
	if err != nil {
		panic(err)
	}
	d, err := ProccessImage(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}

func createFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

// The mime type of the image is changed, it is compressed and then saved in the specified folder.
func ProccessImage(buffer []byte) (string, error) {
	createFolder("screenshots")
	currentTime := time.Now()

	screenName := fmt.Sprintf("%d-%d-%d_%d:%d:%d", currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second())

	filename := screenName + ".webp"

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: 90})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(fmt.Sprintf("./screenshots/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}
