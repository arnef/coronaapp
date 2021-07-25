package viewmodel

import (
	"strings"

	"github.com/arnef/coronaapp/app/covpass"
	"github.com/arnef/coronaapp/app/storage/euvaluerepo"
)

func recoveryData(recovery *covpass.Recovery, rows []*DataRow) []*DataRow {
	return append(rows, []*DataRow{
		{
			Title:    "Krankheit oder Erreger, von dem die Bürgerin oder der Bürger genesen ist / Disease or agent the citizen has recovered from",
			Subtitle: euvaluerepo.GetDiseaseAgentName(recovery.TargetDisease),
		},
		{
			Title:    "Datum des ersten positiven Testergebnisses / Date of first positive test result (YYYY-MM-DD)",
			Subtitle: recovery.FirstResult,
		},
		{
			Title:    "Land der Testung / Member State of test",
			Subtitle: euvaluerepo.GetCountryName(recovery.Country),
		},
		{
			Title:    "Zertifikataussteller / Certificate issuer",
			Subtitle: recovery.CertificateIssuer,
		},
		{
			Title:    "Zertifikat gültig ab / Certificate valid from (YYYY-MM-DD)",
			Subtitle: recovery.ValidFrom,
		},
		{
			Title:    "Zertifikat gültig bis (höchstens 180 Tage ab dem Datum des ersten positiven Testergebnisses) / Certificate valid until (no more than 180 days after the date of first positive test result)(YYYY-MM-DD)",
			Subtitle: recovery.ValidUntil,
		},
		{
			Title:    "Zertifikatkennung / Unique certificate identifier",
			Subtitle: strings.TrimPrefix(recovery.ID, "URN:UVCI:"),
		},
	}...)
}
