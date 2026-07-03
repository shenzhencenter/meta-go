package enums

type CustomaudienceUseForProducts string

const (
	CustomaudienceUseForProductsAds               CustomaudienceUseForProducts = "ADS"
	CustomaudienceUseForProductsMarketingMessages CustomaudienceUseForProducts = "MARKETING_MESSAGES"
)

func (value CustomaudienceUseForProducts) String() string {
	return string(value)
}
