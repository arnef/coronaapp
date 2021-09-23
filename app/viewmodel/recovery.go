package viewmodel

import (
	"github.com/arnef/covcert/pkg/covcert"
	"github.com/leonelquinteros/gotext"
)

func recoveryData(recovery covcert.RecoveryCert, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent the citizen has recovered from"), "Disease or agent the citizen has recovered from"),
			Subtitle: recovery.TargetDisease(),
		},
		{
			Title:    intl(gotext.Get("Date of first positive test result (YYYY-MM-DD)"), "Date of first positive test result (YYYY-MM-DD)"),
			Subtitle: recovery.DateOfFirstPositiveTest(),
		},
		{
			Title:    intl(gotext.Get("Member State of test"), "Member State of test"),
			Subtitle: recovery.MemberStateOfTest(),
		},
		{
			Title:    intl(gotext.Get("Certificate valid from (YYYY-MM-DD)"), "Certificate valid from (YYYY-MM-DD)"),
			Subtitle: recovery.ValidFrom().Format("2006-01-02"),
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: recovery.UniqueIdentifier(),
		},
	}...)
}
