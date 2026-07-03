package enums

type ProductcatalogdiagnosticgroupSeverity string

const (
	ProductcatalogdiagnosticgroupSeverityMustFix     ProductcatalogdiagnosticgroupSeverity = "MUST_FIX"
	ProductcatalogdiagnosticgroupSeverityOpportunity ProductcatalogdiagnosticgroupSeverity = "OPPORTUNITY"
)

func (value ProductcatalogdiagnosticgroupSeverity) String() string {
	return string(value)
}
