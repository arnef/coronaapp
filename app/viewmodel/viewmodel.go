package viewmodel

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/arnef/coronaapp/app/utils"
	"github.com/stapelberg/coronaqr"
)

type CertList struct {
	Size  int
	certs []*Cert
}

func (c *CertList) Get(index int) *Cert {
	if index < len(c.certs) {
		return c.certs[index]
	}
	return nil
}

func (c *CertList) CertID(cert *coronaqr.Decoded) (string, error) {
	data, err := json.Marshal(cert)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha1.Sum(data)), nil
}

func (c *CertList) RemoveByID(id string) {
	idx := -1
	for i, c := range c.certs {
		if c.ID == id {
			idx = i
			break
		}
	}
	if idx > -1 {
		c.certs = append(c.certs[:idx], c.certs[idx+1:]...)
		c.Size = len(c.certs)
	}
}

func (c *CertList) Append(cert *utils.CoronaCert) {
	if !c.inList(cert.ID) {

		viewmodelCert := &Cert{
			ID:          cert.ID,
			FamilyName:  cert.Name.FamilyName,
			GivenName:   cert.Name.GivenName,
			DateOfBirth: cert.BirthDateFormatted(),
			vaccination: cert.Vaccination(),
			recovery:    cert.Recovery(),
			test:        cert.Test(),
		}
		viewmodelCert.GenerateData()
		c.certs = append(c.certs, viewmodelCert)
		c.Size = len(c.certs)
	}
}

func (c *CertList) inList(id string) bool {
	for _, c := range c.certs {
		if c.ID == id {
			return true
		}
	}
	return false
}

type VaccinationCertList struct {
	Size  int
	certs []*VaccinationCert
}

func (v *VaccinationCertList) Get(index int) *VaccinationCert {
	if index < len(v.certs) {
		return v.certs[index]
	}
	return nil
}

type VaccinationCert struct {
	VaccinatedOn   string
	Doses          int
	DoseSeries     int
	Target         string
	MedicalProduct string
	Vaccine        string
	Manufacturer   string
	Country        string
	Issuer         string
	CertificateID  string
}

type RecoveryCert struct {
	Target        string
	Country       string
	CertificateID string
	Issuer        string
}

type RecoveryCertList struct {
	Size  int
	certs []*RecoveryCert
}

func (t *RecoveryCertList) Get(index int) *RecoveryCert {
	if index < len(t.certs) {
		return t.certs[index]
	}
	return nil
}
