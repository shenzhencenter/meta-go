package enums

type ApplicationadnetworkanalyticsOrderingTypeEnumParam string

const (
	ApplicationadnetworkanalyticsOrderingTypeEnumParamAscending  ApplicationadnetworkanalyticsOrderingTypeEnumParam = "ASCENDING"
	ApplicationadnetworkanalyticsOrderingTypeEnumParamDescending ApplicationadnetworkanalyticsOrderingTypeEnumParam = "DESCENDING"
)

func (value ApplicationadnetworkanalyticsOrderingTypeEnumParam) String() string {
	return string(value)
}
