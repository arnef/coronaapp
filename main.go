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

	"github.com/arnef/coronaapp/app"
	"github.com/arnef/coronaapp/app/provider"
	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/coronaapp/app/utils"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

func main() {
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

	context.SetVar("scanner", &scanner)
	context.SetVar("myapp", &state)

	win := component.CreateWindow(nil)

	state.Root = win.Root()

	scanner.Win = win
	scanner.Scanned = func(s string) {
		log.Debugln("scanned", s)
		cert, err := utils.CertFromString(s)

		if err != nil {
			log.Errorln(err)
		} else {
			// only allow vaccination certs for now
			if len(cert.Cert.VaccineRecords) > 0 {
				storage.WriteFile(fmt.Sprintf("%s.pem", cert.ID), []byte(cert.Raw))
				state.AppendCert(cert)
			}
		}

	}

	win.Show()
	win.Wait()

	return nil
}

type QRScanner struct {
	Root      qml.Object
	Win       *qml.Window
	HasResult bool
	reader    gozxing.Reader
	Scanned   func(string)
}

func (qr *QRScanner) Scan() {
	log.Debugln("should scan")
	if qr.Win != nil {
		go func() {
			log.Debugln("scan screen for qr")
			img := qr.Win.Snapshot()

			// prepare BinaryBitmap
			bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

			// decode image
			result, _ := qr.reader.Decode(bmp, nil)
			if result != nil {
				log.Debugln(result.GetBarcodeFormat())
				qr.HasResult = true
				qml.Changed(qr, &qr.HasResult)
				qr.Scanned(result.String())
				qr.HasResult = false
			}
		}()
	}
}
