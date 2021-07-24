package covpass

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/minvws/base45-go/eubase45"
)

func DecodeString(val string) (*CovCert, error) {
	unprefixed, err := unprefix(val)
	if err != nil {
		return nil, err
	}
	compressed, err := base45decode(unprefixed)
	if err != nil {
		return nil, err
	}
	coseData, err := decompress(compressed)
	if err != nil {
		return nil, err
	}

	unverified, err := decodeCOSE(coseData)
	if err != nil {
		return nil, err
	}

	cert := unverified.claims.HCert.DCC
	cert.ValidFrom = time.Unix(unverified.claims.Iat, 0)
	cert.ValidUntil = time.Unix(unverified.claims.Exp, 0)

	return &cert, nil
}

func DecodeBytes(val []byte) (*CovCert, error) {
	return DecodeString(string(val))
}

type decoded struct {
	cert       CovCert
	issuedAt   time.Time
	expiration time.Time
}

func unprefix(prefixObject string) (string, error) {
	if !strings.HasPrefix(prefixObject, "HC1:") {
		return "", errors.New("data does not start with HC1: prefix")
	}

	return strings.TrimPrefix(prefixObject, "HC1:"), nil
}

func base45decode(encoded string) ([]byte, error) {
	return eubase45.EUBase45Decode([]byte(encoded))
}

func decompress(compressed []byte) ([]byte, error) {
	zr, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, zr); err != nil {
		return nil, err
	}
	if err := zr.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeCOSE(coseData []byte) (*unverifiedCOSE, error) {
	var v signedCWT
	if err := cbor.Unmarshal(coseData, &v); err != nil {
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
type hcert struct {
	DCC CovCert `cbor:"1,keyasint"`
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

type signedCWT struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected map[interface{}]interface{}
	Payload     []byte
	Signature   []byte
}

type unverifiedCOSE struct {
	v      signedCWT
	p      coseHeader
	claims claims
}
