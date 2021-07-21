package euvaluerepo

import (
	"github.com/pariz/gountries"
)

func GetDiseaseValue(key string) string {
	if val, ok := diseaseMap[key]; ok {
		return val
	}
	return key
}

func GetManufacturerValue(key string) string {
	if val, ok := manufacturerMap[key]; ok {
		return val
	}
	return key
}

func GetMedicalValue(key string) string {
	if val, ok := medicalMap[key]; ok {
		return val
	}
	return key
}

func GetVaccineProphylaxisValue(key string) string {
	if val, ok := vaccineMap[key]; ok {
		return val
	}
	return key
}

func GetCountryValue(key string) string {
	query := gountries.New()

	if country, err := query.FindCountryByAlpha(key); err == nil {
		return country.Name.Common
	}

	return key
}
