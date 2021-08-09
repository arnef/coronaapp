package viewmodel

import (
	"fmt"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
	"github.com/leonelquinteros/gotext"
)

func vaccinationData(vaccination *covpass.Vaccination, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent targeted"), "Disease or agent targeted"),
			Subtitle: euvaluerepo.GetDiseaseAgentName(vaccination.TargetDisease),
		},
		{
			Title:    intl(gotext.Get("Vaccine"), "Vaccine"),
			Subtitle: euvaluerepo.GetProductName(vaccination.Product),
		},
		{
			Title:    intl(gotext.Get("Vaccine Type"), "Vaccine Type"),
			Subtitle: euvaluerepo.GetProphylaxisName(vaccination.VaccineCode),
		},
		{
			Title:    intl(gotext.Get("Manufacturer"), "Manufacturer"),
			Subtitle: euvaluerepo.GetManufacturerName(vaccination.Manufacturer),
		},
		{
			Title:    intl(gotext.Get("Number in a series of vaccinations/doses"), "Number in a series of vaccinations/doses"),
			Subtitle: fmt.Sprintf("%d/%d", vaccination.DoseNumber, vaccination.TotalSerialDoses),
		},
		{
			Title:    intl(gotext.Get("Date of vaccination (YYYY-MM-DD)"), "Date of vaccination (YYYY-MM-DD)"),
			Subtitle: vaccination.Occurence,
		},
		{
			Title:    intl(gotext.Get("Member State of vaccination"), "Member State of vaccination"),
			Subtitle: euvaluerepo.GetCountryName(vaccination.Country),
		},
		{
			Title:    intl(gotext.Get("Certificate issuer"), "Certificate issuer"),
			Subtitle: vaccination.CertificateIssuer,
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: CleanID(vaccination.ID),
		},
	}...)
}
