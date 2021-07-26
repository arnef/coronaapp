/// values copied from
/// https://github.com/Digitaler-Impfnachweis/covpass-android/blob/3b95b05221e8108ea8cc3cb8e4268a5c1af124a6/covpass-sdk/src/main/java/de/rki/covpass/sdk/storage/EUValueSetRepository.kt
package euvaluerepo

var medicalMap = map[string]string{
	"EU/1/20/1528":                     "Comirnaty",
	"EU/1/20/1507":                     "COVID-19 Vaccine Moderna",
	"EU/1/21/1529":                     "Vaxzevria",
	"EU/1/20/1525":                     "COVID-19 Vaccine Janssen",
	"CVnCoV":                           "CVnCoV",
	"Sputnik-V":                        "Sputnik-V",
	"Convidecia":                       "Convidecia",
	"EpiVacCorona":                     "EpiVacCorona",
	"BBIBP-CorV":                       "BBIBP-CorV",
	"Inactivated-SARS-CoV-2-Vero-Cell": "Inactivated SARS-CoV-2 (Vero Cell)",
	"CoronaVac":                        "CoronaVac",
	"Covaxin":                          "Covaxin (also known as BBV152 A, B, C)",
}

var diseaseMap = map[string]string{
	"840539006": "COVID-19",
}

var manufacturerMap = map[string]string{
	"ORG-100001699":               "AstraZeneca AB",
	"ORG-100030215":               "Biontech Manufacturing GmbH",
	"ORG-100031184":               "Moderna Biotech Spain S.L.",
	"ORG-100006270":               "Curevac AG",
	"ORG-100013793":               "CanSino Biologics",
	"ORG-100020693":               "China Sinopharm International Corp. - Beijing location",
	"ORG-100010771":               "Sinopharm Weiqida Europe Pharmaceutical s.r.o. - Prague location",
	"ORG-100024420":               "Sinopharm Zhijun (Shenzhen) Pharmaceutical Co. Ltd. - Shenzhen location",
	"ORG-100032020":               "Novavax CZ AS",
	"Gamaleya-Research-Institute": "Gamaleya Research Institute",
	"Vector-Institute":            "Vector Institute",
	"Sinovac-Biotech":             "Sinovac Biotech",
	"Bharat-Biotech":              "Bharat Biotech",
}

var vaccineMap = map[string]string{
	"1119349007": "SARS-CoV-2 mRNA vaccine",
	"1119305005": "SARS-CoV-2 antigen vaccine",
	"J07BX03":    "covid-19 vaccines",
}

var testTypeMap = map[string]string{
	"LP6464-4":   "Nucleic acid amplification with probe detection",
	"LP217198-3": "Rapid immunoassay",
}

var testManufacturerMap = map[string]string{
	"1232": "Abbott Rapid Diagnostics, Panbio COVID-19 Ag Test",
	"1304": "AMEDA Labordiagnostik GmbH, AMP Rapid Test SARS-CoV-2 Ag",
	"1065": "Becton Dickinson, Veritor System Rapid Detection of SARS-CoV-2",
	"1331": "Beijing Lepu Medical Technology Co., Ltd, SARS-CoV-2 Antigen Rapid Test Kit",
	"1484": "Beijing Wantai Biological Pharmacy Enterprise Co., Ltd, Wantai SARS-CoV-2 Ag Rapid Test (FIA)",
	"1242": "Bionote, Inc, NowCheck COVID-19 Ag Test",
	"1223": "BIOSYNEX SWISS SA, BIOSYNEX COVID-19 Ag BSS",
	"1173": "CerTest Biotec, S.L., CerTest SARS-CoV-2 Card test",
	"1244": "GenBody, Inc, Genbody COVID-19 Ag Test",
	"1360": "Guangdong Wesail Biotech Co., Ltd, COVID-19 Ag Test Kit",
	"1363": "Hangzhou Clongene Biotech Co., Ltd, Covid-19 Antigen Rapid Test Kit",
	"1767": "Healgen Scientific Limited Liability Company, Coronavirus Ag Rapid Test Cassette",
	"1333": "Joinstar Biomedical Technology Co., Ltd, COVID-19 Rapid Antigen Test (Colloidal Gold)",
	"1268": "LumiraDX UK Ltd, LumiraDx SARS-CoV-2 Ag Test",
	"1180": "MEDsan GmbH, MEDsan SARS-CoV-2 Antigen Rapid Test",
	"1481": "MP Biomedicals Germany GmbH, Rapid SARS-CoV-2 Antigen Test Card",
	"1162": "Nal von minden GmbH, NADAL COVID-19 Ag Test",
	"1271": "Precision Biosensor, Inc, Exdia COVID-19 Ag",
	"1341": "Qingdao Hightop Biotech Co., Ltd, SARS-CoV-2 Antigen Rapid Test (Immunochromatography)",
	"1097": "Quidel Corporation, Sofia SARS Antigen FIA",
	"1489": "Safecare Biotech (Hangzhou) Co. Ltd, COVID-19 Antigen Rapid Test Kit (Swab)",
	"344":  "SD BIOSENSOR Inc, STANDARD F COVID-19 Ag FIA",
	"345":  "SD BIOSENSOR Inc, STANDARD Q COVID-19 Ag Test",
	"1218": "Siemens Healthineers, CLINITEST Rapid Covid-19 Antigen Test",
	"1278": "Xiamen Boson Biotech Co. Ltd, Rapid SARS-CoV-2 Antigen Test Card",
	"1343": "Zhejiang Orient Gene Biotech, Coronavirus Ag Rapid Test Cassette (Swab)",
}

var testResultMap = map[string]string{
	"260415000": "Not detected",
	"260373001": "Detected",
}
