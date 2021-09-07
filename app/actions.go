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

func (s *State) AppendAndPersist(cert string) {
	c, err := utils.CertFromString(cert)
	if err != nil {
		return
	}
	storage.WriteFile(fmt.Sprintf("%s.pem", c.ID), []byte(c.Raw))
	s.AppendCert(c)
}
