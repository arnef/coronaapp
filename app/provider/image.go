package provider

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path"

	"github.com/arnef/coronaapp/app/storage"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	log "github.com/sirupsen/logrus"
)

var cache map[string]image.Image = map[string]image.Image{}

func ImageProvider(imgID string, width, height int) image.Image {
	if width == 0 && height == 0 {
		width = 512
		height = 512
	}
	if cache == nil {
		cache = map[string]image.Image{}
	}
	key := fmt.Sprintf("%s_%dx%d", imgID, width, height)
	if _, exists := cache[key]; !exists {
		writer := qrcode.NewQRCodeWriter()
		name := fmt.Sprintf("%s.pem", imgID)
		config, err := storage.UserConfigDir()
		if err != nil {
			return ErrorCert(err)
		}
		data, err := ioutil.ReadFile(path.Join(config, name))
		if err != nil {
			return ErrorCert(err)
		}
		img, err := writer.Encode(string(data), gozxing.BarcodeFormat_QR_CODE, width, height, nil)
		if err != nil {
			return ErrorCert(err)
		}
		cache[key] = img
	}

	return cache[key]
}

func ErrorCert(err error) image.Image {
	log.Error(err)
	cat, err := os.Open("assets/bad_certs.png")
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(cat)
	if err != nil {
		panic(err)
	}
	cat.Close()
	return img
}
