package viewmodel

import (
	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
)

func testData(test *covpass.Test, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    "Zielkrankheit oder -erreger / Disease or agent targeted",
			Subtitle: euvaluerepo.GetDiseaseAgentName(test.TargetDisease),
		},
		{
			Title:    "Art des Tests / Type of test",
			Subtitle: euvaluerepo.GetTestTypeName(test.TestType),
		},
		{
			Title:    "Produktname / Test name",
			Subtitle: test.TestName,
		},
		{
			Title:    "Testhersteller / Test manufacturer",
			Subtitle: euvaluerepo.GetTestManufacturerName(test.Manufacturer),
		},
		{
			Title:    "Datum und Uhrzeit der Probenahme / Date and time of the sample collection (YYYY-MM-DD, HH:MM)",
			Subtitle: test.SampleCollection.Format("2006-02-01 15:04"),
		},
		{
			Title:    "Testergebnis / Test result",
			Subtitle: euvaluerepo.GetTestResultName(test.TestResult),
		},
		{
			Title:    "Testzentrum oder -einrichtung / Testing centre or facility",
			Subtitle: test.TestCentre,
		},
		{
			Title:    "Land der Testung / Member State of test",
			Subtitle: euvaluerepo.GetCountryName(test.Country),
		},
		{
			Title:    "Zertifikataussteller / Certificate issuer",
			Subtitle: test.CertificateIssuer,
		},
		{
			Title:    "Zertifikatkennung / Unique certificate identifier",
			Subtitle: test.ID,
		},
	}...)
}
