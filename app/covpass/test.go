package covpass

import "time"

const (
	PositivePCRTestCertType     CertType = 5
	NegativePCRTestCertType     CertType = 6
	PositiveAntigenTestCertType CertType = 7
	NegativeAntigenTestCertType CertType = 8
)

type Test struct {
	TargetDisease     string    `cbor:"tg"`
	TestType          string    `cbor:"tt"`
	TestName          string    `cbor:"nm"`
	Manufacturer      string    `cbor:"ma"`
	SampleCollection  time.Time `cbor:"sc"`
	TestResult        string    `cbor:"tr"`
	TestCentre        string    `cbor:"tc"`
	Country           string    `cbor:"co"`
	CertificateIssuer string    `cbor:"is"`
	ID                string    `cbor:"ci"`
}

func (t *Test) Type() CertType {
	switch t.TestType {

	case PCRTest:
		switch t.TestResult {
		case PositiveResult:
			return PositivePCRTestCertType
		case NegativeResult:
			return NegativePCRTestCertType
		}

	case AntigenTest:
		switch t.TestResult {
		case PositiveResult:
			return PositiveAntigenTestCertType
		case NegativeResult:
			return NegativeAntigenTestCertType
		}
	}
	return PositivePCRTestCertType
}

func (t *Test) IsPositive() bool {
	return t.TestResult == PositiveResult
}

const (
	PCRTest        string = "LP6464-4"
	AntigenTest    string = "LP217198-3"
	PositiveResult string = "260373001"
	NegativeResult string = "260415000"
)

const (
	PCRTestExpiryTime     time.Duration = 72 * time.Hour
	AntigenTestExpiryTime time.Duration = 48 * time.Hour
)
