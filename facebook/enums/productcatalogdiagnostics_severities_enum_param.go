package enums

type ProductcatalogdiagnosticsSeveritiesEnumParam string

const (
	ProductcatalogdiagnosticsSeveritiesEnumParamMustFix     ProductcatalogdiagnosticsSeveritiesEnumParam = "MUST_FIX"
	ProductcatalogdiagnosticsSeveritiesEnumParamOpportunity ProductcatalogdiagnosticsSeveritiesEnumParam = "OPPORTUNITY"
)

func (value ProductcatalogdiagnosticsSeveritiesEnumParam) String() string {
	return string(value)
}
