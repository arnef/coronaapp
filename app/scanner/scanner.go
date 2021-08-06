package scanner

import (
	"image"
	"sync"
	"time"

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
	mutex           *sync.Mutex
	result          chan (string)
	onResultHandler OnResultHandler
	HasResult       bool
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

type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func (s *scanner) Scan(x, y, width, height int) {

	go func() {
		log.Debugln("scanner.Scan")
		// lock while reader is decoding
		s.mutex.Lock()
		defer s.mutex.Unlock()
		img, ok := s.win.Snapshot().(SubImage)
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
				s.Decode(result.String())
			}
		}

	}()
}

func (s *scanner) Decode(val string) {
	log.Debugln("scanner.Handle", val)
	time.Sleep(100 * time.Millisecond)
	s.result <- val
}
