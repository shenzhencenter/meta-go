package enums

type PagenlpConfigsModelEnumParam string

const (
	PagenlpConfigsModelEnumParamArabic          PagenlpConfigsModelEnumParam = "ARABIC"
	PagenlpConfigsModelEnumParamChinese         PagenlpConfigsModelEnumParam = "CHINESE"
	PagenlpConfigsModelEnumParamCroatian        PagenlpConfigsModelEnumParam = "CROATIAN"
	PagenlpConfigsModelEnumParamCustom          PagenlpConfigsModelEnumParam = "CUSTOM"
	PagenlpConfigsModelEnumParamDanish          PagenlpConfigsModelEnumParam = "DANISH"
	PagenlpConfigsModelEnumParamDutch           PagenlpConfigsModelEnumParam = "DUTCH"
	PagenlpConfigsModelEnumParamEnglish         PagenlpConfigsModelEnumParam = "ENGLISH"
	PagenlpConfigsModelEnumParamFrenchStandard  PagenlpConfigsModelEnumParam = "FRENCH_STANDARD"
	PagenlpConfigsModelEnumParamGeorgian        PagenlpConfigsModelEnumParam = "GEORGIAN"
	PagenlpConfigsModelEnumParamGermanStandard  PagenlpConfigsModelEnumParam = "GERMAN_STANDARD"
	PagenlpConfigsModelEnumParamGreek           PagenlpConfigsModelEnumParam = "GREEK"
	PagenlpConfigsModelEnumParamHebrew          PagenlpConfigsModelEnumParam = "HEBREW"
	PagenlpConfigsModelEnumParamHungarian       PagenlpConfigsModelEnumParam = "HUNGARIAN"
	PagenlpConfigsModelEnumParamIrish           PagenlpConfigsModelEnumParam = "IRISH"
	PagenlpConfigsModelEnumParamItalianStandard PagenlpConfigsModelEnumParam = "ITALIAN_STANDARD"
	PagenlpConfigsModelEnumParamKorean          PagenlpConfigsModelEnumParam = "KOREAN"
	PagenlpConfigsModelEnumParamNorwegianBokmal PagenlpConfigsModelEnumParam = "NORWEGIAN_BOKMAL"
	PagenlpConfigsModelEnumParamPolish          PagenlpConfigsModelEnumParam = "POLISH"
	PagenlpConfigsModelEnumParamPortuguese      PagenlpConfigsModelEnumParam = "PORTUGUESE"
	PagenlpConfigsModelEnumParamRomanian        PagenlpConfigsModelEnumParam = "ROMANIAN"
	PagenlpConfigsModelEnumParamSpanish         PagenlpConfigsModelEnumParam = "SPANISH"
	PagenlpConfigsModelEnumParamSwedish         PagenlpConfigsModelEnumParam = "SWEDISH"
	PagenlpConfigsModelEnumParamVietnamese      PagenlpConfigsModelEnumParam = "VIETNAMESE"
)

func (value PagenlpConfigsModelEnumParam) String() string {
	return string(value)
}
