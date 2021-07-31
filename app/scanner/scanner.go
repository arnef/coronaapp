package scanner

import (
	"image"
	"sync"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

type OnResultHandler = func(string)

type Scanner interface {
	Wait()
	Scan(x, y, width, height int)
	Decode(val string)
}

func New(win *qml.Window, onResultHandler OnResultHandler) Scanner {
	return &scanner{
		win:             win,
		result:          make(chan string),
		mutex:           &sync.Mutex{},
		onResultHandler: onResultHandler,
	}
}

type scanner struct {
	Root            qml.Object
	win             *qml.Window
	result          chan (string)
	HasResult       bool
	mutex           *sync.Mutex
	onResultHandler OnResultHandler
	Decoding        bool
}

func (s *scanner) Wait() {
	log.Debugln("scanner.Wait")
	s.HasResult = false

	go func() {
		data := <-s.result
		s.HasResult = true
		qml.Changed(s, &s.HasResult)
		if s.onResultHandler != nil {
			s.onResultHandler(data)
		}
		log.Debugln("scanner.Done")
	}()
}

func (s *scanner) Scan(x, y, width, height int) {

	go func() {
		log.Debugln("scanner.Scan")
		// lock while reader is decoding
		s.mutex.Lock()
		defer s.mutex.Unlock()
		img := s.win.Snapshot()
		subImg := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(x, y, x+width, y+height))
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
			s.Decode(result.String())
		}

	}()
}

func (s *scanner) Decode(val string) {
	log.Debugln("scanner.Handle", val)
	s.result <- val
}
