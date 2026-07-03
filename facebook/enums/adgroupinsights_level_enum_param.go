package enums

type AdgroupinsightsLevelEnumParam string

const (
	AdgroupinsightsLevelEnumParamAccount  AdgroupinsightsLevelEnumParam = "account"
	AdgroupinsightsLevelEnumParamAd       AdgroupinsightsLevelEnumParam = "ad"
	AdgroupinsightsLevelEnumParamAdset    AdgroupinsightsLevelEnumParam = "adset"
	AdgroupinsightsLevelEnumParamCampaign AdgroupinsightsLevelEnumParam = "campaign"
)

func (value AdgroupinsightsLevelEnumParam) String() string {
	return string(value)
}
