package enums

type AdaccountactivitiesCategoryEnumParam string

const (
	AdaccountactivitiesCategoryEnumParamAccount    AdaccountactivitiesCategoryEnumParam = "ACCOUNT"
	AdaccountactivitiesCategoryEnumParamAd         AdaccountactivitiesCategoryEnumParam = "AD"
	AdaccountactivitiesCategoryEnumParamAdKeywords AdaccountactivitiesCategoryEnumParam = "AD_KEYWORDS"
	AdaccountactivitiesCategoryEnumParamAdSet      AdaccountactivitiesCategoryEnumParam = "AD_SET"
	AdaccountactivitiesCategoryEnumParamAudience   AdaccountactivitiesCategoryEnumParam = "AUDIENCE"
	AdaccountactivitiesCategoryEnumParamBid        AdaccountactivitiesCategoryEnumParam = "BID"
	AdaccountactivitiesCategoryEnumParamBudget     AdaccountactivitiesCategoryEnumParam = "BUDGET"
	AdaccountactivitiesCategoryEnumParamCampaign   AdaccountactivitiesCategoryEnumParam = "CAMPAIGN"
	AdaccountactivitiesCategoryEnumParamDate       AdaccountactivitiesCategoryEnumParam = "DATE"
	AdaccountactivitiesCategoryEnumParamStatus     AdaccountactivitiesCategoryEnumParam = "STATUS"
	AdaccountactivitiesCategoryEnumParamTargeting  AdaccountactivitiesCategoryEnumParam = "TARGETING"
)

func (value AdaccountactivitiesCategoryEnumParam) String() string {
	return string(value)
}
