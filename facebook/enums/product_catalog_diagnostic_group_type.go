package enums

type ProductcatalogdiagnosticgroupType string

const (
	ProductcatalogdiagnosticgroupTypeArVisibilityIssues            ProductcatalogdiagnosticgroupType = "AR_VISIBILITY_ISSUES"
	ProductcatalogdiagnosticgroupTypeAttributesInvalid             ProductcatalogdiagnosticgroupType = "ATTRIBUTES_INVALID"
	ProductcatalogdiagnosticgroupTypeAttributesMissing             ProductcatalogdiagnosticgroupType = "ATTRIBUTES_MISSING"
	ProductcatalogdiagnosticgroupTypeCategory                      ProductcatalogdiagnosticgroupType = "CATEGORY"
	ProductcatalogdiagnosticgroupTypeCheckout                      ProductcatalogdiagnosticgroupType = "CHECKOUT"
	ProductcatalogdiagnosticgroupTypeDaVisibilityIssues            ProductcatalogdiagnosticgroupType = "DA_VISIBILITY_ISSUES"
	ProductcatalogdiagnosticgroupTypeEventSourceIssues             ProductcatalogdiagnosticgroupType = "EVENT_SOURCE_ISSUES"
	ProductcatalogdiagnosticgroupTypeImageQuality                  ProductcatalogdiagnosticgroupType = "IMAGE_QUALITY"
	ProductcatalogdiagnosticgroupTypeLowQualityTitleAndDescription ProductcatalogdiagnosticgroupType = "LOW_QUALITY_TITLE_AND_DESCRIPTION"
	ProductcatalogdiagnosticgroupTypePolicyViolation               ProductcatalogdiagnosticgroupType = "POLICY_VIOLATION"
	ProductcatalogdiagnosticgroupTypeShopsVisibilityIssues         ProductcatalogdiagnosticgroupType = "SHOPS_VISIBILITY_ISSUES"
)

func (value ProductcatalogdiagnosticgroupType) String() string {
	return string(value)
}
