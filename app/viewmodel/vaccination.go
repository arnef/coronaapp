package viewmodel

import (
	"fmt"
	"strings"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
)

func vaccinationData(vaccination *covpass.Vaccination, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    "Zielkrankheit oder -erreger / Disease or agent targeted",
			Subtitle: euvaluerepo.GetDiseaseAgentName(vaccination.TargetDisease),
		},
		{
			Title:    "Impfstoff / Vaccine",
			Subtitle: euvaluerepo.GetProductName(vaccination.Product),
		},
		{
			Title:    "Art des Impfstoffs / Vaccine Type",
			Subtitle: euvaluerepo.GetProphylaxisName(vaccination.VaccineCode),
		},
		{
			Title:    "Hersteller / Manufacturer",
			Subtitle: euvaluerepo.GetManufacturerName(vaccination.Manufacturer),
		},
		{
			Title:    "Nummer der Impfung / Number in a series of vaccinations/doses",
			Subtitle: fmt.Sprintf("%d/%d", vaccination.DoseNumber, vaccination.TotalSerialDoses),
		},
		{
			Title:    "Datum der Impfung / Date of vaccination (YYYY-MM-DD)",
			Subtitle: vaccination.Occurence,
		},
		{
			Title:    "Land der Impfung / Member State of vaccination",
			Subtitle: euvaluerepo.GetCountryName(vaccination.Country),
		},
		{
			Title:    "Zertifikataussteller / Certificate issuer",
			Subtitle: vaccination.CertificateIssuer,
		},
		{
			Title:    "Zertifikatkennung / Unique certificate identifier",
			Subtitle: strings.TrimPrefix(vaccination.ID, "URN:UVCI:"),
		},
	}...)
}
