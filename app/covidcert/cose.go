package covidcert

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type signedCWT struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected map[interface{}]interface{}
	Payload     []byte
	Signature   []byte
}

type coseHeader struct {
	// Cryptographic algorithm. See COSE Algorithms Registry:
	// https://www.iana.org/assignments/cose/cose.xhtml
	Alg int `cbor:"1,keyasint,omitempty"`
	// Key identifier
	Kid []byte `cbor:"4,keyasint,omitempty"`
	// Full Initialization Vector
	IV []byte `cbor:"5,keyasint,omitempty"`
}
type hcert struct {
	DCC Cert `cbor:"1,keyasint"`
}

type claims struct {
	Iss   string `cbor:"1,keyasint"`
	Sub   string `cbor:"2,keyasint"`
	Aud   string `cbor:"3,keyasint"`
	Exp   int64  `cbor:"4,keyasint"`
	Nbf   int    `cbor:"5,keyasint"`
	Iat   int64  `cbor:"6,keyasint"`
	Cti   []byte `cbor:"7,keyasint"`
	HCert hcert  `cbor:"-260,keyasint"`
}
type unverifiedCOSE struct {
	v      signedCWT
	p      coseHeader
	claims claims
}

func decodeCOSE(data []byte) (*unverifiedCOSE, error) {
	var v signedCWT
	if err := cbor.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal: %v", err)
	}

	var p coseHeader
	if len(v.Protected) > 0 {
		if err := cbor.Unmarshal(v.Protected, &p); err != nil {
			return nil, fmt.Errorf("cbor.Unmarshal(v.Protected): %v", err)
		}
	}

	var c claims
	if err := cbor.Unmarshal(v.Payload, &c); err != nil {
		return nil, fmt.Errorf("cbor.Unmarshal(v.Payload): %v", err)
	}

	return &unverifiedCOSE{
		v:      v,
		p:      p,
		claims: c,
	}, nil
}
