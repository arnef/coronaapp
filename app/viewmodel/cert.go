package viewmodel

import (
	"fmt"
	"time"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/leonelquinteros/gotext"
)

type Cert struct {
	ID          string
	FamilyName  string
	GivenName   string
	DateOfBirth string
	vaccination *covpass.Vaccination
	recovery    *covpass.Recovery
	test        *covpass.Test
	Data        *DataRows
}

func intl(val string, key string) string {
	if val != key {
		return fmt.Sprintf("%s / %s", val, key)
	}
	return val
}

func (c *Cert) ValidUntil() time.Time {
	validUntil := time.Now()

	if c.recovery != nil {
		if vu, err := covpass.ParseDay(c.recovery.ValidUntil); err == nil {

			validUntil = vu
		}
	} else if c.test != nil {
		if c.test.Type() == covpass.NegativePCRTestCertType {
			validUntil = c.test.SampleCollection.AddDate(0, 0, 2)
		} else if c.test.Type() == covpass.NegativeAntigenTestCertType {
			validUntil = c.test.SampleCollection.AddDate(0, 0, 1)
		}
	} else if c.vaccination != nil {
		if occ, err := covpass.ParseDay(c.vaccination.Occurence); err == nil {
			if c.vaccination.Type() == covpass.VaccinationFullProtectionCertType {

				validUntil = occ.AddDate(1, 0, 0)
			} else {
				validUntil = occ.AddDate(0, 0, 14)
			}
		}
	}

	return validUntil
}

func (c *Cert) GenerateData() {
	rows := []*DataRow{
		{
			Title:    intl(gotext.Get("Name, first name"), "Name, first name"),
			Subtitle: fmt.Sprintf("%s, %s", c.FamilyName, c.GivenName),
		},
		{
			Title:    intl(gotext.Get("Date of birth (YYYY-MM-DD)"), "Date of birth (YYYY-MM-DD)"),
			Subtitle: c.DateOfBirth,
		},
	}

	if c.vaccination != nil {
		rows = vaccinationData(c.vaccination, rows)
	}
	if c.recovery != nil {
		rows = recoveryData(c.recovery, rows)
	}
	if c.test != nil {
		rows = testData(c.test, rows)
	}

	c.Data = &DataRows{
		Size: len(rows),
		Rows: rows,
	}
}

func (c *Cert) Type() string {
	if c.vaccination != nil {
		return gotext.Get("Vaccination certificate")
	}
	if c.recovery != nil {
		return gotext.Get("Recovery certificate")
	}

	if c.test != nil {
		switch c.test.TestType {
		case covpass.PCRTest:
			return gotext.Get("PCR test")
		case covpass.AntigenTest:
			return gotext.Get("Rapid antigen test")
		}

	}
	return ""
}

func (c *Cert) Title() string {
	if c.vaccination != nil {
		switch c.vaccination.Type() {
		case covpass.VaccinationFullProtectionCertType:
			return gotext.Get("Full vaccination protection")
		case covpass.VaccinationCompleteCertType:
			occurence, err := covpass.ParseDay(c.vaccination.Occurence)
			if err == nil {
				return gotext.Get("Full protection as of %s", occurence.Add(15*24*time.Hour).Format("02.01.2006"))
			}
		default:
			return gotext.Get("Incomplete vaccination protection")
		}
	}
	if c.recovery != nil {
		validDay, err := covpass.ParseDay(c.recovery.ValidUntil)
		if err == nil {
			return gotext.Get("Valid until %s", validDay.Format("02.01.2006"))
		}
	}
	if c.test != nil {
		return c.test.SampleCollection.Format("02.01.2006 15:04")
	}
	return ""
}

func (c *Cert) Color() string {
	if c.vaccination != nil {
		if c.vaccination.HasFullProtection() {
			return "#0560c4"
		}
		return "#d2e7fe"
	}
	if c.recovery != nil {
		return "#043268"
	}
	if c.test != nil {
		return "#7d29b7"
	}
	return "white"
}

func (c *Cert) Icon() string {
	if c.vaccination != nil {
		if c.vaccination.HasFullProtection() {
			return "cert_complete.svg"
		}
		return "cert_incomplete.svg"
	}
	if c.recovery != nil {
		return "cert_recovery.svg"
	}
	if c.test != nil {
		return "cert_test.svg"
	}

	// fallback
	return "logo.svg"
}
