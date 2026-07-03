package enums

type BusinessadnetworkanalyticsOrderingTypeEnumParam string

const (
	BusinessadnetworkanalyticsOrderingTypeEnumParamAscending  BusinessadnetworkanalyticsOrderingTypeEnumParam = "ASCENDING"
	BusinessadnetworkanalyticsOrderingTypeEnumParamDescending BusinessadnetworkanalyticsOrderingTypeEnumParam = "DESCENDING"
)

func (value BusinessadnetworkanalyticsOrderingTypeEnumParam) String() string {
	return string(value)
}
