package enums

type AdcampaigngroupSpecialAdCategories string

const (
	AdcampaigngroupSpecialAdCategoriesCredit                    AdcampaigngroupSpecialAdCategories = "CREDIT"
	AdcampaigngroupSpecialAdCategoriesEmployment                AdcampaigngroupSpecialAdCategories = "EMPLOYMENT"
	AdcampaigngroupSpecialAdCategoriesFinancialProductsServices AdcampaigngroupSpecialAdCategories = "FINANCIAL_PRODUCTS_SERVICES"
	AdcampaigngroupSpecialAdCategoriesHousing                   AdcampaigngroupSpecialAdCategories = "HOUSING"
	AdcampaigngroupSpecialAdCategoriesIssuesElectionsPolitics   AdcampaigngroupSpecialAdCategories = "ISSUES_ELECTIONS_POLITICS"
	AdcampaigngroupSpecialAdCategoriesNone                      AdcampaigngroupSpecialAdCategories = "NONE"
	AdcampaigngroupSpecialAdCategoriesOnlineGamblingAndGaming   AdcampaigngroupSpecialAdCategories = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdcampaigngroupSpecialAdCategories) String() string {
	return string(value)
}
