package enums

type AdaccountadcreativesAuthorizationCategoryEnumParam string

const (
	AdaccountadcreativesAuthorizationCategoryEnumParamNone                               AdaccountadcreativesAuthorizationCategoryEnumParam = "NONE"
	AdaccountadcreativesAuthorizationCategoryEnumParamPolitical                          AdaccountadcreativesAuthorizationCategoryEnumParam = "POLITICAL"
	AdaccountadcreativesAuthorizationCategoryEnumParamPoliticalWithDigitallyCreatedMedia AdaccountadcreativesAuthorizationCategoryEnumParam = "POLITICAL_WITH_DIGITALLY_CREATED_MEDIA"
)

func (value AdaccountadcreativesAuthorizationCategoryEnumParam) String() string {
	return string(value)
}
