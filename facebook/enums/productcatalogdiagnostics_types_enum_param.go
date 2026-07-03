package enums

type ProductcatalogdiagnosticsTypesEnumParam string

const (
	ProductcatalogdiagnosticsTypesEnumParamArVisibilityIssues            ProductcatalogdiagnosticsTypesEnumParam = "AR_VISIBILITY_ISSUES"
	ProductcatalogdiagnosticsTypesEnumParamAttributesInvalid             ProductcatalogdiagnosticsTypesEnumParam = "ATTRIBUTES_INVALID"
	ProductcatalogdiagnosticsTypesEnumParamAttributesMissing             ProductcatalogdiagnosticsTypesEnumParam = "ATTRIBUTES_MISSING"
	ProductcatalogdiagnosticsTypesEnumParamCategory                      ProductcatalogdiagnosticsTypesEnumParam = "CATEGORY"
	ProductcatalogdiagnosticsTypesEnumParamCheckout                      ProductcatalogdiagnosticsTypesEnumParam = "CHECKOUT"
	ProductcatalogdiagnosticsTypesEnumParamDaVisibilityIssues            ProductcatalogdiagnosticsTypesEnumParam = "DA_VISIBILITY_ISSUES"
	ProductcatalogdiagnosticsTypesEnumParamEventSourceIssues             ProductcatalogdiagnosticsTypesEnumParam = "EVENT_SOURCE_ISSUES"
	ProductcatalogdiagnosticsTypesEnumParamImageQuality                  ProductcatalogdiagnosticsTypesEnumParam = "IMAGE_QUALITY"
	ProductcatalogdiagnosticsTypesEnumParamLowQualityTitleAndDescription ProductcatalogdiagnosticsTypesEnumParam = "LOW_QUALITY_TITLE_AND_DESCRIPTION"
	ProductcatalogdiagnosticsTypesEnumParamPolicyViolation               ProductcatalogdiagnosticsTypesEnumParam = "POLICY_VIOLATION"
	ProductcatalogdiagnosticsTypesEnumParamShopsVisibilityIssues         ProductcatalogdiagnosticsTypesEnumParam = "SHOPS_VISIBILITY_ISSUES"
)

func (value ProductcatalogdiagnosticsTypesEnumParam) String() string {
	return string(value)
}
