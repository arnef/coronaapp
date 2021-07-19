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
	"encoding/json"
	"fmt"
	"log"

	"github.com/arnef/coronaapp/app/covidcert"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/nanu-c/qml-go"
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

	scanner := QRScanner{}
	context := engine.Context()
	context.SetVar("testvar", &scanner)
	// testvar.GetMessage()

	win := component.CreateWindow(nil)

	scanner.Root = win.Root()
	scanner.Win = win
	// testvar.Root = win.Root()
	win.Show()
	win.Wait()

	return nil
}

type QRScanner struct {
	Root      qml.Object
	Win       *qml.Window
	HasResult bool
	Result    string
}

func (qr *QRScanner) Scan() {
	fmt.Println("I Shoud scan")
	if qr.Win != nil {

		fmt.Println("scan screen for qr")
		img := qr.Win.Snapshot()

		// prepare BinaryBitmap
		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

		// decode image
		qrReader := qrcode.NewQRCodeReader()
		result, _ := qrReader.Decode(bmp, nil)
		if result != nil {
			cert, err := covidcert.Decode(result.String())
			if err == nil {
				certJson, err := json.Marshal(cert)
				if err == nil {
					qr.HasResult = true
					qr.Result = string(certJson)
					qml.Changed(qr, &qr.HasResult)
					qml.Changed(qr, &qr.Result)
				}
			}
		}
	}
}
