package enums

type AdaccountactivitiesDataSourceEnumParam string

const (
	AdaccountactivitiesDataSourceEnumParamCalypso      AdaccountactivitiesDataSourceEnumParam = "CALYPSO"
	AdaccountactivitiesDataSourceEnumParamTao          AdaccountactivitiesDataSourceEnumParam = "TAO"
	AdaccountactivitiesDataSourceEnumParamTaoAdAccount AdaccountactivitiesDataSourceEnumParam = "TAO_AD_ACCOUNT"
	AdaccountactivitiesDataSourceEnumParamTaoAdStatus  AdaccountactivitiesDataSourceEnumParam = "TAO_AD_STATUS"
)

func (value AdaccountactivitiesDataSourceEnumParam) String() string {
	return string(value)
}
