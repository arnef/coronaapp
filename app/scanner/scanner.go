package scanner

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

type Scanner interface {
	Scan(x, y, width, height int)
}

func New(win *qml.Window) Scanner {
	return &scanner{
		win:    win,
		result: nil,
	}
}

type scanner struct {
	Root   qml.Object
	win    *qml.Window
	result chan string
	Text   string
}
type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func (s *scanner) prepareImage(subImg image.Image) image.Image {
	subImg = imaging.Grayscale(subImg)
	subImg = imaging.AdjustSigmoid(subImg, .5, 10)
	subImg = imaging.AdjustFunc(subImg, func(c color.NRGBA) color.NRGBA {
		if c.R > 128 {
			return color.NRGBA{
				R: 255, G: 255, B: 255, A: 255,
			}
		}
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	})
	return subImg
}

func (s *scanner) working(x, y, width, height int, img SubImage) {
	subImg := img.SubImage(image.Rect(x, y, x+width, y+height))
	if width > 512 || height > 512 {
		subImg = imaging.Resize(subImg, 512, 512, imaging.Lanczos)
	}
	subImg = s.prepareImage(subImg)
	bmp, err := gozxing.NewBinaryBitmapFromImage(subImg)
	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bmp, map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
	})
	if err != nil {
		log.Error(err)
		return
	}
	if result != nil && s.result != nil {
		s.result <- result.String()
	}
}

func (s *scanner) Scan(x, y, width, height int) {
	if s.result == nil {
		s.result = make(chan string)
		go s.wait()
	}
	log.Debugln("scanner.Scan", x, y, width, height)
	if s.win == nil {
		log.Debug("No window Skip Scan")
		return
	}
	img, ok := s.win.Snapshot().(SubImage)
	if ok {
		go s.working(x, y, width, height, img)
	}
}

func (s *scanner) wait() {
	s.Text = <-s.result
	qml.Changed(s, &s.Text)
	s.result = nil
}
