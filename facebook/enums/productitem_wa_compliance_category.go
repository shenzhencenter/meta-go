package enums

type ProductitemWaComplianceCategory string

const (
	ProductitemWaComplianceCategoryCountryOriginExempt ProductitemWaComplianceCategory = "COUNTRY_ORIGIN_EXEMPT"
	ProductitemWaComplianceCategoryDefault             ProductitemWaComplianceCategory = "DEFAULT"
)

func (value ProductitemWaComplianceCategory) String() string {
	return string(value)
}
