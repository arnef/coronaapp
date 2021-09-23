package viewmodel

import (
	"github.com/arnef/covcert/pkg/covcert"
	"github.com/leonelquinteros/gotext"
)

func testData(test covcert.TestCert, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent targeted"), "Disease or agent targeted"),
			Subtitle: test.TargetDisease(),
		},
		{
			Title:    intl(gotext.Get("Type of test"), "Type of test"),
			Subtitle: test.TypeOfTest(),
		},
		// {
		// 	Title:    intl(gotext.Get("Test name"), "Test name"),
		// 	Subtitle: test,,
		// },
		// {
		// 	Title:    intl(gotext.Get("Test manufacturer"), "Test manufacturer"),
		// 	Subtitle: euvaluerepo.GetTestManufacturerName(test.Manufacturer),
		// },
		{
			Title:    intl(gotext.Get("Date and time of the sample collection (YYYY-MM-DD, HH:MM)"), "Date and time of the sample collection (YYYY-MM-DD, HH:MM)"),
			Subtitle: test.DateTimeOfSampleCollection().Format("2006-01-02, 15:04"),
		},
		{
			Title:    intl(gotext.Get("Test result"), "Test result"),
			Subtitle: test.TestResult(),
		},
		{
			Title:    intl(gotext.Get("Testing centre or facility"), "Testing centre or facility"),
			Subtitle: test.TestingCentre(),
		},
		{
			Title:    intl(gotext.Get("Member State of test"), "Member State of test"),
			Subtitle: test.MemberStateOfTest(),
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: test.UniqueIdentifier(),
		},
	}...)
}
