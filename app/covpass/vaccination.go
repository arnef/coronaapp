package covpass

type Vaccination struct {
	TargetDisease     string `cbor:"tg"`
	VaccineCode       string `cbor:"vp"`
	Product           string `cbor:"mp"`
	Manufacturer      string `cbor:"ma"`
	DoseNumber        int    `cbor:"dn"`
	TotalSerialDoses  int    `cbor:"sd"`
	Occurence         string `cbor:"dt"`
	Country           string `cbor:"co"`
	CertificateIssuer string `cbor:"is"`
	ID                string `json:"ci"`
}

func (v *Vaccination) IsComplete() bool {
	return v.DoseNumber == v.TotalSerialDoses
}

func (v *Vaccination) IsCompleteSingleDose() bool {
	return v.DoseNumber == 1 && v.TotalSerialDoses == 1
}

func (v *Vaccination) HasFullProtection() bool {

	return v.IsComplete() && v.DaysSinceOccurence() > 14
}

func (v *Vaccination) DaysSinceOccurence() int {
	occurence, err := ParseDay(v.Occurence)
	if err != nil {
		panic(err)
	}
	return int(Today().Sub(occurence).Hours()) / 24
}

func (v *Vaccination) Type() CertType {
	if v.HasFullProtection() {
		return VaccinationFullProtectionCertType
	} else if v.IsComplete() {
		return VaccinationCompleteCertType
	}
	return VaccinationIncompleteCertType
}

type CertType int

const (
	VaccinationFullProtectionCertType CertType = 0
	VaccinationCompleteCertType       CertType = 1
	VaccinationIncompleteCertType     CertType = 2
)
