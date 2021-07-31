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
	"os"

	"github.com/arnef/coronaapp/app"
	"github.com/arnef/coronaapp/app/provider"
	"github.com/arnef/coronaapp/app/scanner"
	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/coronaapp/app/utils"
	"github.com/leonelquinteros/gotext"
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
	context := engine.Context()

	r := R{
		Delete:     gotext.Get("Delete"),
		DeleteCert: gotext.Get("Delete certificate?"),
		Cancel:     gotext.Get("Cancel"),
	}
	context.SetVar("myapp", &state)
	context.SetVar("R", &r)

	win := component.CreateWindow(nil)

	scanner := scanner.New(win, func(s string) {
		cert, err := utils.CertFromString(s)
		if err != nil {
			// TODO display error message
			log.Error(err)
			return
		}
		go storage.WriteFile(fmt.Sprintf("%s.pem", cert.ID), []byte(cert.Raw))
		state.AppendCert(cert)
	})
	context.SetVar("scanner", scanner)
	state.Root = win.Root()

	win.Show()
	win.Wait()

	return nil
}

type R struct {
	Delete     string
	DeleteCert string
	Cancel     string
}
