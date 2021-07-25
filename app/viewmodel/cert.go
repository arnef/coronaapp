package viewmodel

import (
	"fmt"

	"github.com/arnef/coronaapp/app/covpass"
)

type Cert struct {
	ID          string
	FamilyName  string
	GivenName   string
	DateOfBirth string
	Data        *DataRows
	vaccination *covpass.Vaccination
	recovery    *covpass.Recovery
	test        *covpass.Test
}

func (c *Cert) GenerateData() {
	rows := []*DataRow{
		{
			Title:    "Name, Vorname / Name, first name",
			Subtitle: fmt.Sprintf("%s, %s", c.FamilyName, c.GivenName),
		},
		{
			Title:    "Geburtstag / Date of birth (YYYY-MM-DD)",
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
		return "Impfzertifikat"
	}
	if c.recovery != nil {
		return "Genesenenzertifikat"
	}

	if c.test != nil {
		switch c.test.TestType {
		case covpass.PCRTest:
			return "PCR-Test"
		case covpass.AntigenTest:
			return "Antigen-Schnelltest"
		}

	}
	return ""
}

func (c *Cert) Title() string {
	if c.vaccination != nil {
		switch c.vaccination.Type() {
		case covpass.VaccinationFullProtectionCertType:
			return "Vollständiger Impfschutz"
		case covpass.VaccinationCompleteCertType:
			return "Vollständig ab ..."
		case covpass.VaccinationIncompleteCertType:
			return "Unvollständiger Impfschutz"
		}
	}
	if c.recovery != nil {
		validDay, err := covpass.ParseDay(c.recovery.ValidUntil)
		if err == nil {
			return fmt.Sprintf("Gültig bis %s", validDay.Format("02.01.2006"))
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
