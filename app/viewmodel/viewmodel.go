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
			// VaccinationCerts: &VaccinationCertList{
			// 	Size:  len(cert.Cert.VaccineRecords),
			// 	certs: make([]*VaccinationCert, len(cert.Cert.VaccineRecords)),
			// },
			// RecoveryCerts: &RecoveryCertList{
			// 	Size:  len(cert.Cert.RecoveryRecords),
			// 	certs: make([]*RecoveryCert, len(cert.Cert.RecoveryRecords)),
			// },
		}

		// for i, vac := range cert.Cert.VaccineRecords {

		// 	viewmodelCert.VaccinationCerts.certs[i] = &VaccinationCert{
		// 		VaccinatedOn:   vac.Date,
		// 		Doses:          int(vac.Doses),
		// 		DoseSeries:     int(vac.DoseSeries),
		// 		Target:         euvaluerepo.GetDiseaseValue(vac.Target),
		// 		MedicalProduct: euvaluerepo.GetMedicalValue(vac.Product),
		// 		Vaccine:        euvaluerepo.GetVaccineProphylaxisValue(vac.Vaccine),
		// 		Manufacturer:   euvaluerepo.GetManufacturerValue(vac.Manufacturer),
		// 		Country:        euvaluerepo.GetCountryValue(vac.Country),
		// 		Issuer:         vac.Issuer,
		// 		CertificateID:  strings.Replace(vac.CertificateID, "URN:UVCI:", "", 1),
		// 	}
		// }

		// for i, rec := range cert.Cert.RecoveryRecords {
		// 	viewmodelCert.RecoveryCerts.certs[i] = &RecoveryCert{
		// 		Target:        euvaluerepo.GetDiseaseValue(rec.Target),
		// 		Issuer:        rec.Issuer,
		// 		Country:       euvaluerepo.GetCountryValue(rec.Country),
		// 		CertificateID: rec.CertificateID,
		// 	}
		// }

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
