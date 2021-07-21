package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/stapelberg/coronaqr"
)

type CoronaCert struct {
	*coronaqr.Decoded
	Raw string
	ID  string
}

func CertFromString(val string) (*CoronaCert, error) {
	decoded, err := coronaqr.Decode(val)
	if err != nil {
		return nil, err
	}
	return &CoronaCert{
		Decoded: decoded.SkipVerification(),
		ID:      fmt.Sprintf("%x", sha1.Sum([]byte(val))),
		Raw:     val,
	}, nil
}
