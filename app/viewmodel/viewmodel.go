package viewmodel

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
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

func (c *CertList) Append(cert *utils.CoronaCert) {
	if !c.inList(cert.ID) {

		viewmodelCert := &Cert{
			ID:          cert.ID,
			FamilyName:  cert.Cert.PersonalName.FamilyName,
			GivenName:   cert.Cert.PersonalName.GivenName,
			DateOfBirth: cert.Cert.DateOfBirth,
			VaccinationCerts: &VaccinationCertList{
				Size:  len(cert.Cert.VaccineRecords),
				certs: make([]*VaccinationCert, len(cert.Cert.VaccineRecords)),
			},
		}

		for i, vac := range cert.Cert.VaccineRecords {

			viewmodelCert.VaccinationCerts.certs[i] = &VaccinationCert{
				VaccinatedOn:   vac.Date,
				Doses:          int(vac.Doses),
				DoseSeries:     int(vac.DoseSeries),
				Target:         euvaluerepo.GetDiseaseValue(vac.Target),
				MedicalProduct: euvaluerepo.GetMedicalValue(vac.Product),
				Vaccine:        euvaluerepo.GetVaccineProphylaxisValue(vac.Vaccine),
				Manufacturer:   euvaluerepo.GetManufacturerValue(vac.Manufacturer),
				Country:        euvaluerepo.GetCountryValue(vac.Country),
				Issuer:         vac.Issuer,
				CertificateID:  strings.Replace(vac.CertificateID, "URN:UVCI:", "", 1),
			}
		}

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

type Cert struct {
	ID               string
	FamilyName       string
	GivenName        string
	DateOfBirth      string
	VaccinationCerts *VaccinationCertList
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
