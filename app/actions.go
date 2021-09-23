package app

import (
	"crypto/sha1"
	"fmt"

	"github.com/arnef/coronaapp/app/storage"
	"github.com/arnef/covcert/pkg/covcert"
	"github.com/arnef/covcert/pkg/decoder"
	"github.com/nanu-c/qml-go"
)

func (s *State) AppendCert(cert covcert.CovCert) {
	s.Certs.Append(cert)
	qml.Changed(s, &s.Certs)
}

func (s *State) RemoveCert(id string) {
	s.Certs.RemoveByID(id)
	storage.RmFile(fmt.Sprintf("%s.pem", id))
	qml.Changed(s, &s.Certs)
}

func (s *State) AppendAndPersist(cert string) {
	// c, err := utils.CertFromString(cert)
	c, err := decoder.DecodeString(cert)
	if err != nil {
		return
	}
	val := []byte(cert)
	storage.WriteFile(fmt.Sprintf("%s.pem", fmt.Sprintf("%x", sha1.Sum(val))), val)
	s.AppendCert(c)
}
