package enums

type AdcampaigninsightsLevelEnumParam string

const (
	AdcampaigninsightsLevelEnumParamAccount  AdcampaigninsightsLevelEnumParam = "account"
	AdcampaigninsightsLevelEnumParamAd       AdcampaigninsightsLevelEnumParam = "ad"
	AdcampaigninsightsLevelEnumParamAdset    AdcampaigninsightsLevelEnumParam = "adset"
	AdcampaigninsightsLevelEnumParamCampaign AdcampaigninsightsLevelEnumParam = "campaign"
)

func (value AdcampaigninsightsLevelEnumParam) String() string {
	return string(value)
}
