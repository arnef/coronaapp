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
	"os"

	"github.com/arnef/coronaapp/app"
	"github.com/arnef/coronaapp/app/provider"
	"github.com/arnef/coronaapp/app/scanner"
	"github.com/arnef/coronaapp/app/storage"
	"github.com/leonelquinteros/gotext"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

func main() {
	// TODO set by build envs
	// log.SetLevel(log.DebugLevel)

	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gotext.Configure("./locales", os.Getenv("LANGUAGE"), "default")
	r := R{
		Delete:     gotext.Get("Delete"),
		DeleteCert: gotext.Get("Delete certificate?"),
		Cancel:     gotext.Get("Cancel"),
	}

	engine := qml.NewEngine()
	engine.AddImageProvider(storage.AppName, provider.ImageProvider)

	state := app.Init()
	context := engine.Context()

	component, err := engine.LoadFile("qml/Main.qml")
	if err != nil {
		return err
	}

	context.SetVar("myapp", &state)
	context.SetVar("R", &r)

	win := component.CreateWindow(nil)
	scanner := scanner.New(win)

	context.SetVar("scanner", scanner)

	win.Show()
	win.Wait()

	return nil
}

type R struct {
	Delete     string
	DeleteCert string
	Cancel     string
}
