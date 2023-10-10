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

var folderName = "./screenshots/"
var pngFile string
var Screename string

// Decode the base64 srtring sent from js callback to png file
func DecodeImage(b64 string, domain string) {
	// Create folder for screenshots and check if exist or not
	createFolder(folderName)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	pngFilename := domain + ".png"
	f, err := os.Create(folderName + pngFilename)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	ff, _ := os.Open(folderName + pngFilename)
	buffer, err := io.ReadAll(ff)
	if err != nil {
		panic(err)
	}
	// Start the proccess of optmizing the png file to reduce the size
	d, err := ProccessImage(buffer)
	if err != nil {
		panic(err)
	}
	// clean png file after finish the optmiztion
	cleanAfterOptmize(folderName + pngFilename)

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
	fmt.Println(Screename)
	currentTime := time.Now()

	// create the uniqu id for the new screen filename
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

	// convert the file and edit quality
	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: 90})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(folderName+filename, processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}

func cleanAfterOptmize(fileName string) {
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}
}
