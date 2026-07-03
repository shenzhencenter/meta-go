package enums

type AdaccountinsightsLevelEnumParam string

const (
	AdaccountinsightsLevelEnumParamAccount  AdaccountinsightsLevelEnumParam = "account"
	AdaccountinsightsLevelEnumParamAd       AdaccountinsightsLevelEnumParam = "ad"
	AdaccountinsightsLevelEnumParamAdset    AdaccountinsightsLevelEnumParam = "adset"
	AdaccountinsightsLevelEnumParamCampaign AdaccountinsightsLevelEnumParam = "campaign"
)

func (value AdaccountinsightsLevelEnumParam) String() string {
	return string(value)
}
