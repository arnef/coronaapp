package covpass

import (
	"fmt"
	"time"
)

type CovCert struct {
	Issuer       string        `cbor:"-"`
	ValidFrom    time.Time     `cbor:"-"`
	ValidUntil   time.Time     `cbor:"-"`
	Version      string        `cbor:"ver"`
	Name         Name          `cbor:"nam"`
	DateOfBirth  string        `cbor:"dob"`
	Vaccinations []Vaccination `cbor:"v"`
	Recoveries   []Recovery    `cbor:"r"`
	Tests        []Test        `cbor:"t"`
}

func (c *CovCert) Vaccination() *Vaccination {
	if len(c.Vaccinations) > 0 {
		return &c.Vaccinations[0]
	}
	return nil
}

func (c *CovCert) Recovery() *Recovery {
	if len(c.Recoveries) > 0 {
		return &c.Recoveries[0]
	}
	return nil
}

func (c *CovCert) Test() *Test {
	if len(c.Tests) > 0 {
		return &c.Tests[0]
	}
	return nil
}

func (c *CovCert) FullName() string {
	return fmt.Sprintf("%s %s", c.Name.GivenName, c.Name.FamilyName)
}

func (c *CovCert) FullNameReverse() string {
	return fmt.Sprintf("%s, %s", c.Name.FamilyName, c.Name.GivenName)
}

func (c *CovCert) BirthDateFormatted() string {
	switch len(c.DateOfBirth) {
	case 0: // empty
		return "XXXX-XX-XX"
	case 4: // "2021"
		return fmt.Sprintf("%s-XX-XX", c.DateOfBirth)
	case 7: // "2021-01"
		return fmt.Sprintf("%s-XX", c.DateOfBirth)
	default:
		return c.DateOfBirth
	}
}

type Name struct {
	FamilyName    string `cbor:"fn"`
	FamilyNameStd string `cbor:"fnt"`
	GivenName     string `cbor:"gn"`
	GivenNameStd  string `cbor:"gnt"`
}
