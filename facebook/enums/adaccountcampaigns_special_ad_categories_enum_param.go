package enums

type AdaccountcampaignsSpecialAdCategoriesEnumParam string

const (
	AdaccountcampaignsSpecialAdCategoriesEnumParamCredit                    AdaccountcampaignsSpecialAdCategoriesEnumParam = "CREDIT"
	AdaccountcampaignsSpecialAdCategoriesEnumParamEmployment                AdaccountcampaignsSpecialAdCategoriesEnumParam = "EMPLOYMENT"
	AdaccountcampaignsSpecialAdCategoriesEnumParamFinancialProductsServices AdaccountcampaignsSpecialAdCategoriesEnumParam = "FINANCIAL_PRODUCTS_SERVICES"
	AdaccountcampaignsSpecialAdCategoriesEnumParamHousing                   AdaccountcampaignsSpecialAdCategoriesEnumParam = "HOUSING"
	AdaccountcampaignsSpecialAdCategoriesEnumParamIssuesElectionsPolitics   AdaccountcampaignsSpecialAdCategoriesEnumParam = "ISSUES_ELECTIONS_POLITICS"
	AdaccountcampaignsSpecialAdCategoriesEnumParamNone                      AdaccountcampaignsSpecialAdCategoriesEnumParam = "NONE"
	AdaccountcampaignsSpecialAdCategoriesEnumParamOnlineGamblingAndGaming   AdaccountcampaignsSpecialAdCategoriesEnumParam = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdaccountcampaignsSpecialAdCategoriesEnumParam) String() string {
	return string(value)
}
