package enums

type AdcampaignactivitiesCategoryEnumParam string

const (
	AdcampaignactivitiesCategoryEnumParamAccount    AdcampaignactivitiesCategoryEnumParam = "ACCOUNT"
	AdcampaignactivitiesCategoryEnumParamAd         AdcampaignactivitiesCategoryEnumParam = "AD"
	AdcampaignactivitiesCategoryEnumParamAdKeywords AdcampaignactivitiesCategoryEnumParam = "AD_KEYWORDS"
	AdcampaignactivitiesCategoryEnumParamAdSet      AdcampaignactivitiesCategoryEnumParam = "AD_SET"
	AdcampaignactivitiesCategoryEnumParamAudience   AdcampaignactivitiesCategoryEnumParam = "AUDIENCE"
	AdcampaignactivitiesCategoryEnumParamBid        AdcampaignactivitiesCategoryEnumParam = "BID"
	AdcampaignactivitiesCategoryEnumParamBudget     AdcampaignactivitiesCategoryEnumParam = "BUDGET"
	AdcampaignactivitiesCategoryEnumParamCampaign   AdcampaignactivitiesCategoryEnumParam = "CAMPAIGN"
	AdcampaignactivitiesCategoryEnumParamDate       AdcampaignactivitiesCategoryEnumParam = "DATE"
	AdcampaignactivitiesCategoryEnumParamStatus     AdcampaignactivitiesCategoryEnumParam = "STATUS"
	AdcampaignactivitiesCategoryEnumParamTargeting  AdcampaignactivitiesCategoryEnumParam = "TARGETING"
)

func (value AdcampaignactivitiesCategoryEnumParam) String() string {
	return string(value)
}
