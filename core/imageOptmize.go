package core

import (
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"time"

	"image/png"
	_ "image/png"
	"os"
	"strings"

	"github.com/h2non/bimg"
)

var folderName = "./screenshots/"

func Optmize(base string, domain string) (string, error) {
	// Create folder for screenshots and check if exist or not
	err := createFolder(folderName)
	if err != nil {
		return "", err
	}

	err2 := decodeImage(base, domain)

	if err2 != nil {
		return "", err2
	}
	fileName, _ := proccessImage(folderName + domain + ".png")

	// clean png file after finish the optmiztion
	cleanAfterOptmize(folderName + domain + ".png")

	return fileName, nil
}

// Decode the base64 srtring sent from js callback to png file
func decodeImage(b64 string, domain string) error {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	m, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	pngFilename := domain + ".png"
	f, err := os.Create(folderName + pngFilename)
	if err != nil {
		return err
	}

	err = png.Encode(f, m)
	if err != nil {
		return err
	}

	return nil
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
func proccessImage(pngFilename string) (string, error) {
	ff, _ := os.Open(pngFilename)
	buffer, err := io.ReadAll(ff)
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	// Start the proccess of optmizing the png file to reduce the size
	currentTime := time.Now()
	// create the uniqu id for the new screen filename
	screenName := fmt.Sprintf("%s_%d-%d-%d_%d:%d:%d", strings.Replace(pngFilename, folderName, "", 20), currentTime.Year(),
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
