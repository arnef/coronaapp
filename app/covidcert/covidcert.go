package covidcert

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/minvws/base45-go/eubase45"
)

type CovidCert struct {
	Cert       Cert      `json:"cert"`
	IssuedAt   time.Time `json:"issued_at"`
	Expiration time.Time `json:"expiration"`
}

type Cert struct {
	Version        string          `cbor:"ver" json:"ver"`
	PersonalName   Name            `cbor:"nam" json:"name"`
	DateOfBirth    string          `cbor:"dob" json:"dob"`
	VaccineRecords []VaccineRecord `cbor:"v" json:"v"`
}

type Name struct {
	FamilyName    string `cbor:"fn" json:"fn"`
	FamilyNameStd string `cbor:"fnt" json:"fnt"`
	GivenName     string `cbor:"gn" json:"gn"`
	GivenNameStd  string `cbor:"gnt" json:"gnt"`
}

type VaccineRecord struct {
	Target        string `cbor:"tg" json:"tg"`
	Vaccine       string `cbor:"vp" json:"vp"`
	Product       string `cbor:"mp" json:"mp"`
	Manufacturer  string `cbor:"ma" json:"ma"`
	Doses         int    `cbor:"dn" json:"dn"`
	DoseSeries    int    `cbor:"sd" json:"sd"`
	Date          string `cbor:"dt" json:"dt"`
	Country       string `cbor:"co" json:"co"`
	Issuer        string `cbor:"is" json:"is"`
	CertificateID string `cbor:"ci" json:"ci"`
}

func Decode(data string) (*CovidCert, error) {
	if !strings.HasPrefix(data, "HC1:") {
		return nil, fmt.Errorf("not a valid covid cert string")
	}

	compressed, err := base45decode(data[4:])
	if err != nil {
		return nil, err
	}
	coseData, err := decompress(compressed)
	if err != nil {
		return nil, err
	}
	certData, err := decodeCOSE(coseData)

	return &CovidCert{
		Cert:       certData.claims.HCert.DCC,
		IssuedAt:   time.Unix(certData.claims.Iat, 0),
		Expiration: time.Unix(certData.claims.Exp, 0),
	}, err
}

func base45decode(val string) ([]byte, error) {
	return eubase45.EUBase45Decode([]byte(val))
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
	return buf.Bytes(), nil
}
