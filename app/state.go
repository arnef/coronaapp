package app

import (
	"strings"

	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/coronaapp/app/viewmodel"
	"github.com/arnef/covcert/pkg/decoder"
	"github.com/leonelquinteros/gotext"
	"github.com/nanu-c/qml-go"

	log "github.com/sirupsen/logrus"
)

type State struct {
	Root  qml.Object
	Certs *viewmodel.CertList
}

func Init() State {
	state := State{
		Certs: &viewmodel.CertList{
			EmptyMessage: gotext.Get("You currently do not have saved any digital EU-COVID certificates."),
			Size:         0,
		},
	}

	files, err := storage.ReadDir("")

	if err == nil {
		for _, f := range files {
			if strings.HasSuffix(f, ".pem") {
				data, err := storage.ReadFile(f)
				if err == nil {
					// cert, err := utils.CertFromString((string(data)))
					cert, err := decoder.DecodeString(string(data))
					if err != nil {
						log.Error(err)
						continue
					}

					state.Certs.Append(cert)

				}
			}
		}

	}
	return state
}
