package core

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"io"
	"log"
	"os"
	"strings"

	"github.com/h2non/bimg"
)

func DecodeImage(b64 string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	pngFilename := "test.png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	as, _ := os.Open("test.png")
	buffer, err := io.ReadAll(as)
	if err != nil {
		panic(err)
	}

	errDir := createFolder("uploads")
	if errDir != nil {
		panic(errDir)
	}

	w, err := imageProcessing(buffer, 90, "uploads")
	if err != nil {
		panic(err)
	}
	fmt.Println(w)
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
func imageProcessing(buffer []byte, quality int, dirname string) (string, error) {
	filename := strings.Replace("wrwrr", "-", "", -1) + ".webp"

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
	if writeError != nil {
		return filename, writeError
	}

	return filename, nil
}
