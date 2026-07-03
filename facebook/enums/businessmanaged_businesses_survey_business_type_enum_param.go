package enums

type BusinessmanagedBusinessesSurveyBusinessTypeEnumParam string

const (
	BusinessmanagedBusinessesSurveyBusinessTypeEnumParamAdvertiser   BusinessmanagedBusinessesSurveyBusinessTypeEnumParam = "ADVERTISER"
	BusinessmanagedBusinessesSurveyBusinessTypeEnumParamAgency       BusinessmanagedBusinessesSurveyBusinessTypeEnumParam = "AGENCY"
	BusinessmanagedBusinessesSurveyBusinessTypeEnumParamAppDeveloper BusinessmanagedBusinessesSurveyBusinessTypeEnumParam = "APP_DEVELOPER"
	BusinessmanagedBusinessesSurveyBusinessTypeEnumParamPublisher    BusinessmanagedBusinessesSurveyBusinessTypeEnumParam = "PUBLISHER"
)

func (value BusinessmanagedBusinessesSurveyBusinessTypeEnumParam) String() string {
	return string(value)
}
