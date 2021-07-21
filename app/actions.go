package app

import (
	"github.com/arnef/coronaapp/app/utils"
	"github.com/nanu-c/qml-go"
)

func (s *State) AppendCert(cert *utils.CoronaCert) {
	s.Certs.Append(cert)
	qml.Changed(s, &s.Certs)
}
