package viewmodel

import (
	"strings"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
	"github.com/leonelquinteros/gotext"
)

func testData(test *covpass.Test, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent targeted"), "Disease or agent targeted"),
			Subtitle: euvaluerepo.GetDiseaseAgentName(test.TargetDisease),
		},
		{
			Title:    intl(gotext.Get("Type of test"), "Type of test"),
			Subtitle: euvaluerepo.GetTestTypeName(test.TestType),
		},
		{
			Title:    intl(gotext.Get("Test name"), "Test name"),
			Subtitle: test.TestName,
		},
		{
			Title:    intl(gotext.Get("Test manufacturer"), "Test manufacturer"),
			Subtitle: euvaluerepo.GetTestManufacturerName(test.Manufacturer),
		},
		{
			Title:    intl(gotext.Get("Date and time of the sample collection (YYYY-MM-DD, HH:MM)"), "Date and time of the sample collection (YYYY-MM-DD, HH:MM)"),
			Subtitle: test.SampleCollection.Format("2006-02-01, 15:04"),
		},
		{
			Title:    intl(gotext.Get("Test result"), "Test result"),
			Subtitle: euvaluerepo.GetTestResultName(test.TestResult),
		},
		{
			Title:    intl(gotext.Get("Testing centre or facility"), "Testing centre or facility"),
			Subtitle: test.TestCentre,
		},
		{
			Title:    intl(gotext.Get("Member State of test"), "Member State of test"),
			Subtitle: euvaluerepo.GetCountryName(test.Country),
		},
		{
			Title:    intl(gotext.Get("Certificate issuer"), "Certificate issuer"),
			Subtitle: test.CertificateIssuer,
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: strings.TrimPrefix(test.ID, "URN:UVCI:"),
		},
	}...)
}
