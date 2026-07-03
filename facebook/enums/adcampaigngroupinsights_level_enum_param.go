package enums

type AdcampaigngroupinsightsLevelEnumParam string

const (
	AdcampaigngroupinsightsLevelEnumParamAccount  AdcampaigngroupinsightsLevelEnumParam = "account"
	AdcampaigngroupinsightsLevelEnumParamAd       AdcampaigngroupinsightsLevelEnumParam = "ad"
	AdcampaigngroupinsightsLevelEnumParamAdset    AdcampaigngroupinsightsLevelEnumParam = "adset"
	AdcampaigngroupinsightsLevelEnumParamCampaign AdcampaigngroupinsightsLevelEnumParam = "campaign"
)

func (value AdcampaigngroupinsightsLevelEnumParam) String() string {
	return string(value)
}
