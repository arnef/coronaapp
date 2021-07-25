package euvaluerepo

import (
	"github.com/pariz/gountries"
)

func GetDiseaseAgentName(key string) string {
	if val, ok := diseaseMap[key]; ok {
		return val
	}
	return key
}

func GetManufacturerName(key string) string {
	if val, ok := manufacturerMap[key]; ok {
		return val
	}
	return key
}

func GetProductName(key string) string {
	if val, ok := medicalMap[key]; ok {
		return val
	}
	return key
}

func GetProphylaxisName(key string) string {
	if val, ok := vaccineMap[key]; ok {
		return val
	}
	return key
}

func GetCountryName(key string) string {
	query := gountries.New()

	if country, err := query.FindCountryByAlpha(key); err == nil {
		return country.Name.Common
	}

	return key
}

func GetTestTypeName(key string) string {
	if val, ok := testTypeMap[key]; ok {
		return val
	}
	return key
}

func GetTestManufacturerName(key string) string {
	if val, ok := testManufacturerMap[key]; ok {
		return val
	}
	return key
}

func GetTestResultName(key string) string {
	if val, ok := testResultMap[key]; ok {
		return val
	}
	return key
}
