package enums

type BusinessownedBusinessesSurveyBusinessTypeEnumParam string

const (
	BusinessownedBusinessesSurveyBusinessTypeEnumParamAdvertiser   BusinessownedBusinessesSurveyBusinessTypeEnumParam = "ADVERTISER"
	BusinessownedBusinessesSurveyBusinessTypeEnumParamAgency       BusinessownedBusinessesSurveyBusinessTypeEnumParam = "AGENCY"
	BusinessownedBusinessesSurveyBusinessTypeEnumParamAppDeveloper BusinessownedBusinessesSurveyBusinessTypeEnumParam = "APP_DEVELOPER"
	BusinessownedBusinessesSurveyBusinessTypeEnumParamPublisher    BusinessownedBusinessesSurveyBusinessTypeEnumParam = "PUBLISHER"
)

func (value BusinessownedBusinessesSurveyBusinessTypeEnumParam) String() string {
	return string(value)
}
