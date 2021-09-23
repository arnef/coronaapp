package viewmodel

import (
	"fmt"

	"github.com/arnef/covcert/pkg/covcert"
	"github.com/leonelquinteros/gotext"
)

func vaccinationData(vaccination covcert.VaccinationCert, rows []*DataRow) []*DataRow {
	doses, total := vaccination.Vaccinations()
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent targeted"), "Disease or agent targeted"),
			Subtitle: vaccination.TargetDisease(),
		},
		{
			Title:    intl(gotext.Get("Vaccine"), "Vaccine"),
			Subtitle: vaccination.Vaccine(),
		},
		{
			Title:    intl(gotext.Get("Vaccine Type"), "Vaccine Type"),
			Subtitle: vaccination.VaccineType(),
		},
		{
			Title:    intl(gotext.Get("Manufacturer"), "Manufacturer"),
			Subtitle: vaccination.Manufacturer(),
		},
		{
			Title:    intl(gotext.Get("Number in a series of vaccinations/doses"), "Number in a series of vaccinations/doses"),
			Subtitle: fmt.Sprintf("%d/%d", doses, total),
		},
		{
			Title:    intl(gotext.Get("Date of vaccination (YYYY-MM-DD)"), "Date of vaccination (YYYY-MM-DD)"),
			Subtitle: vaccination.DateOfVaccination(),
		},
		{
			Title:    intl(gotext.Get("Member State of vaccination"), "Member State of vaccination"),
			Subtitle: vaccination.MemberStateOfVaccination(),
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: vaccination.UniqueIdentifier(),
		},
	}...)
}
