package enums

type BusinessadCustomDerivedMetricsScopeEnumParam string

const (
	BusinessadCustomDerivedMetricsScopeEnumParamAccount            BusinessadCustomDerivedMetricsScopeEnumParam = "ACCOUNT"
	BusinessadCustomDerivedMetricsScopeEnumParamBusiness           BusinessadCustomDerivedMetricsScopeEnumParam = "BUSINESS"
	BusinessadCustomDerivedMetricsScopeEnumParamBusinessAssetGroup BusinessadCustomDerivedMetricsScopeEnumParam = "BUSINESS_ASSET_GROUP"
)

func (value BusinessadCustomDerivedMetricsScopeEnumParam) String() string {
	return string(value)
}
