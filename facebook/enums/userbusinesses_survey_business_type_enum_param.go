package enums

type UserbusinessesSurveyBusinessTypeEnumParam string

const (
	UserbusinessesSurveyBusinessTypeEnumParamAdvertiser   UserbusinessesSurveyBusinessTypeEnumParam = "ADVERTISER"
	UserbusinessesSurveyBusinessTypeEnumParamAgency       UserbusinessesSurveyBusinessTypeEnumParam = "AGENCY"
	UserbusinessesSurveyBusinessTypeEnumParamAppDeveloper UserbusinessesSurveyBusinessTypeEnumParam = "APP_DEVELOPER"
	UserbusinessesSurveyBusinessTypeEnumParamPublisher    UserbusinessesSurveyBusinessTypeEnumParam = "PUBLISHER"
)

func (value UserbusinessesSurveyBusinessTypeEnumParam) String() string {
	return string(value)
}
