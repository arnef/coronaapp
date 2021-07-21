package app

import (
	"fmt"

	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/coronaapp/app/utils"
	"github.com/nanu-c/qml-go"
)

func (s *State) AppendCert(cert *utils.CoronaCert) {
	s.Certs.Append(cert)
	qml.Changed(s, &s.Certs)
}

func (s *State) RemoveCert(id string) {
	s.Certs.RemoveByID(id)
	storage.RmFile(fmt.Sprintf("%s.pem", id))
	qml.Changed(s, &s.Certs)
}
