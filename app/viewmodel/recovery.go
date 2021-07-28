package viewmodel

import (
	"strings"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
	"github.com/leonelquinteros/gotext"
)

func recoveryData(recovery *covpass.Recovery, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    intl(gotext.Get("Disease or agent the citizen has recovered from"), "Disease or agent the citizen has recovered from"),
			Subtitle: euvaluerepo.GetDiseaseAgentName(recovery.TargetDisease),
		},
		{
			Title:    intl(gotext.Get("Date of first positive test result (YYYY-MM-DD)"), "Date of first positive test result (YYYY-MM-DD)"),
			Subtitle: recovery.FirstResult,
		},
		{
			Title:    intl(gotext.Get("Member State of test"), "Member State of test"),
			Subtitle: euvaluerepo.GetCountryName(recovery.Country),
		},
		{
			Title:    intl(gotext.Get("Certificate issuer"), "Certificate issuer"),
			Subtitle: recovery.CertificateIssuer,
		},
		{
			Title:    intl(gotext.Get("Certificate valid from (YYYY-MM-DD)"), "Certificate valid from (YYYY-MM-DD)"),
			Subtitle: recovery.ValidFrom,
		},
		{
			Title:    intl(gotext.Get("Certificate valid until (no more than 180 days after the date of first positive test result)(YYYY-MM-DD)"), "Certificate valid until (no more than 180 days after the date of first positive test result)(YYYY-MM-DD)"),
			Subtitle: recovery.ValidUntil,
		},
		{
			Title:    intl(gotext.Get("Unique certificate identifier"), "Unique certificate identifier"),
			Subtitle: strings.TrimPrefix(recovery.ID, "URN:UVCI:"),
		},
	}...)
}
