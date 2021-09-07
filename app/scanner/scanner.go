package scanner

import (
	"image"
	"sync"

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
		win:   win,
		mutex: &sync.Mutex{},
	}
}

type scanner struct {
	Root  qml.Object
	win   *qml.Window
	mutex *sync.Mutex
	Text  string
}
type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func (s *scanner) Scan(x, y, width, height int) {

	go func() {
		log.Debugln("scanner.Scan")
		if s.win == nil {
			return
		}
		// lock while reader is decoding
		s.mutex.Lock()
		img, ok := s.win.Snapshot().(SubImage)
		s.mutex.Unlock()
		if ok {
			subImg := img.SubImage(image.Rect(x, y, x+width, y+height))
			reader := qrcode.NewQRCodeReader()
			bmp, err := gozxing.NewBinaryBitmapFromImage(subImg)
			if err != nil {
				log.Error(err)
				return
			}
			result, err := reader.DecodeWithoutHints(bmp)
			if err != nil {
				log.Error(err)
				return
			}
			if result != nil {

				s.done(result.String())
			}
		}

	}()
}

func (s *scanner) done(val string) {
	s.mutex.Lock()
	s.Text = val
	qml.Changed(s, &s.Text)
	s.mutex.Unlock()
}
