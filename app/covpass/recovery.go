package covpass

type Recovery struct {
	TargetDisease     string `cbor:"tg"`
	FirstResult       string `cbor:"fr"`
	ValidFrom         string `cbor:"df"`
	ValidUntil        string `cbor:"du"`
	Country           string `cbor:"co"`
	CertificateIssuer string `cbor:"is"`
	ID                string `cbor:"ci"`
}

func (r *Recovery) Type() CertType {
	return RecoveryCertType
}

const RecoveryCertType CertType = 4
