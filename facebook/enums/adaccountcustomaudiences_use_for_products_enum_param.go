package enums

type AdaccountcustomaudiencesUseForProductsEnumParam string

const (
	AdaccountcustomaudiencesUseForProductsEnumParamAds               AdaccountcustomaudiencesUseForProductsEnumParam = "ADS"
	AdaccountcustomaudiencesUseForProductsEnumParamMarketingMessages AdaccountcustomaudiencesUseForProductsEnumParam = "MARKETING_MESSAGES"
)

func (value AdaccountcustomaudiencesUseForProductsEnumParam) String() string {
	return string(value)
}
