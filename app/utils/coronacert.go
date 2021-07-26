package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/arnef/coronaapp/app/covpass"
)

type CoronaCert struct {
	*covpass.CovCert
	Raw string
	ID  string
}

func CertFromString(val string) (*CoronaCert, error) {
	covCert, err := covpass.DecodeString(val)
	if err != nil {
		return nil, err
	}
	return &CoronaCert{
		CovCert: covCert,
		ID:      fmt.Sprintf("%x", sha1.Sum([]byte(val))),
		Raw:     val,
	}, nil
}
