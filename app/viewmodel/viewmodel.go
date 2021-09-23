package viewmodel

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"time"

	"github.com/arnef/covcert/pkg/covcert"
	"github.com/leonelquinteros/gotext"
)

type CertList struct {
	Size         int
	EmptyMessage string
	certs        []*Cert
}

func (c *CertList) Get(index int) *Cert {
	if index < len(c.certs) {
		return c.certs[index]
	}
	return nil
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

func (c *CertList) Append(cert covcert.CovCert) {
	id := fmt.Sprintf("%x", sha1.Sum([]byte(cert.String())))
	if !c.inList(id) {
		viewmodelCert := &Cert{
			ID:          id,
			FullName:    fmt.Sprintf("%s %s", cert.Holder().GivenName, cert.Holder().FamilyName),
			DateOfBirth: cert.Holder().DateOfBirth,
			ValidUntil:  cert.ValidUntil(),
			Title:       gotext.Get("Certificate"), // default unkown type for nl2
			SubTitle:    gotext.Get("Valid until %s", cert.ValidUntil().Format("02.01.2006")),
			Icon:        "no_certs.svg",
			Color:       "#eeeeee",
			TextColor:   "#000000",
		}
		fill(cert, viewmodelCert)
		viewmodelCert.GenerateData(cert)
		c.certs = append(c.certs, viewmodelCert)
		sort.Slice(c.certs, func(a, b int) bool {
			return c.certs[a].ValidUntil.After(c.certs[b].ValidUntil)
		})

		c.Size = len(c.certs)
	}
}
func (c *Cert) GenerateData(cert covcert.CovCert) {
	rows := []*DataRow{
		{
			Title:    intl(gotext.Get("Name, first name"), "Name, first name"),
			Subtitle: fmt.Sprintf("%s, %s", cert.Holder().FamilyName, cert.Holder().GivenName),
		},
		{
			Title:    intl(gotext.Get("Date of birth (YYYY-MM-DD)"), "Date of birth (YYYY-MM-DD)"),
			Subtitle: c.DateOfBirth,
		},
		{
			Title:    intl(gotext.Get("Certificate issuer"), "Certificate issuer"),
			Subtitle: cert.Issuer(),
		},
		{
			Title:    intl(gotext.Get("Valid until (YYYY-MM-DD)"), "Valid until (YYYY-MM-DD)"),
			Subtitle: cert.ValidUntil().Format("2006-01-02"),
		},
	}

	if v, ok := cert.(covcert.VaccinationCert); ok {
		rows = vaccinationData(v, rows)
	} else if r, ok := cert.(covcert.RecoveryCert); ok {
		rows = recoveryData(r, rows)
	} else if t, ok := cert.(covcert.TestCert); ok {
		rows = testData(t, rows)
	}

	c.Data = &DataRows{
		Size: len(rows),
		Rows: rows,
	}
}
func fill(cert covcert.CovCert, model *Cert) {
	if v, ok := cert.(covcert.VaccinationCert); ok {
		model.Title = gotext.Get("Vaccination certificate")
		model.Color = "#d2e7fe"
		model.SubTitle = gotext.Get("Incomplete vaccination protection")
		model.Icon = "cert_incomplete.svg"
		if v.HasFullProtection(time.Now()) {
			model.Color = "#0560c4"
			model.TextColor = "#ffffff"
			model.SubTitle = gotext.Get("Full vaccination protection")
			model.Icon = "cert_complete.svg"
		} else if v.IsComplete() {
			occ, err := time.Parse("2006-01-02", v.DateOfVaccination())
			if err == nil {
				model.SubTitle = gotext.Get("Full protection as of %s", occ.Add(15*24*time.Hour).Format("02.01.2006"))
			}
		}

	} else if _, ok := cert.(covcert.RecoveryCert); ok {
		model.Title = gotext.Get("Recovery certificate")
		model.SubTitle = gotext.Get("Valid until %s", cert.ValidUntil().Format("02.01.2006"))
		model.Color = "#043268"
		model.TextColor = "#ffffff"
		model.Icon = "cert_recovery.svg"
	} else if t, ok := cert.(covcert.TestCert); ok {
		model.Icon = "cert_test.svg"
		model.Color = "#7d29b7"
		model.TextColor = "#ffffff"
		model.SubTitle = t.DateTimeOfSampleCollection().Format("02.01.2006 15:04")
		// TODO implement test type enums in covcert
		if t.TypeOfTest() == "Rapid immunoassay" {
			model.Title = gotext.Get("Rapid antigen test")
		} else {
			model.Title = gotext.Get("PCR test")
		}
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
