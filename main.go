/*
 * Copyright (C) 2021  arnef
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * coronaapp is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"image"
	"os"

	"github.com/arnef/coronaapp/app"
	"github.com/arnef/coronaapp/app/provider"
	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/coronaapp/app/utils"
	"github.com/leonelquinteros/gotext"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

func main() {
	// TODO set by build envs
	// log.SetLevel(log.DebugLevel)

	gotext.Configure("./locales", os.Getenv("LANGUAGE"), "default")

	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	engine := qml.NewEngine()

	component, err := engine.LoadFile("qml/Main.qml")
	if err != nil {
		return err
	}
	state := app.Init()

	engine.AddImageProvider(storage.AppName, provider.ImageProvider)
	scanner := QRScanner{
		reader: qrcode.NewQRCodeReader(),
	}
	context := engine.Context()

	r := R{
		Delete:     gotext.Get("Delete"),
		DeleteCert: gotext.Get("Delete certificate?"),
		Cancel:     gotext.Get("Cancel"),
	}

	context.SetVar("scanner", &scanner)
	context.SetVar("myapp", &state)
	context.SetVar("R", &r)

	win := component.CreateWindow(nil)

	state.Root = win.Root()

	scanner.Win = win
	scanner.Scanned = func(s string) {
		log.Debugln("scanned", s)
		cert, err := utils.CertFromString(s)

		if err != nil {
			log.Errorln(err)
		} else {
			storage.WriteFile(fmt.Sprintf("%s.pem", cert.ID), []byte(cert.Raw))
			state.AppendCert(cert)
		}

	}

	win.Show()
	win.Wait()

	return nil
}

type R struct {
	Delete     string
	DeleteCert string
	Cancel     string
}

type QRScanner struct {
	Root      qml.Object
	Win       *qml.Window
	HasResult bool
	reader    gozxing.Reader
	Scanned   func(string)
	running   bool
}

func (qr *QRScanner) Scan(x, y, width, height int) {
	if qr.Win != nil {
		go func() {
			if !qr.running {
				log.Debugln("scan screen for qr", x, y, width, height)
				qr.running = true
				img := qr.Win.Snapshot()
				my_sub_image := img.(interface {
					SubImage(r image.Rectangle) image.Image
				}).SubImage(image.Rect(x, y, x+width, y+height))
				// prepare BinaryBitmap
				bmp, _ := gozxing.NewBinaryBitmapFromImage(my_sub_image)
				// decode image
				result, _ := qr.reader.Decode(bmp, nil)
				if result != nil {
					qr.HandleString(result.String())
				}
				qr.running = false
			}
		}()
	}
}

func (qr *QRScanner) HandleString(val string) {
	qr.HasResult = true
	qml.Changed(qr, &qr.HasResult)
	qr.Scanned(val)
	qr.HasResult = false
	qml.Changed(qr, &qr.HasResult)
}
