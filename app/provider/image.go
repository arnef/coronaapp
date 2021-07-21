package provider

import (
	"fmt"
	"image"
	"io/ioutil"
	"path"

	"github.com/arnef/coronaapp/app/storage"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

var cache map[string]image.Image = map[string]image.Image{}

func ImageProvider(imgID string, width, height int) image.Image {
	if width == 0 && height == 0 {
		width = 512
		height = 512
	}
	key := fmt.Sprintf("%s_%dx%d", imgID, width, height)
	if _, exists := cache[key]; !exists {
		writer := qrcode.NewQRCodeWriter()
		name := fmt.Sprintf("%s.pem", imgID)
		config, err := storage.UserConfigDir()
		if err != nil {
			// todo handle errors
			panic(err)
		}
		data, err := ioutil.ReadFile(path.Join(config, name))
		if err != nil {
			panic(err)
		}
		img, _ := writer.Encode(string(data), gozxing.BarcodeFormat_QR_CODE, width, height, nil)
		cache[key] = img
	}

	return cache[key]
}
