package enums

type AdcampaigngroupSpecialAdCategory string

const (
	AdcampaigngroupSpecialAdCategoryCredit                    AdcampaigngroupSpecialAdCategory = "CREDIT"
	AdcampaigngroupSpecialAdCategoryEmployment                AdcampaigngroupSpecialAdCategory = "EMPLOYMENT"
	AdcampaigngroupSpecialAdCategoryFinancialProductsServices AdcampaigngroupSpecialAdCategory = "FINANCIAL_PRODUCTS_SERVICES"
	AdcampaigngroupSpecialAdCategoryHousing                   AdcampaigngroupSpecialAdCategory = "HOUSING"
	AdcampaigngroupSpecialAdCategoryIssuesElectionsPolitics   AdcampaigngroupSpecialAdCategory = "ISSUES_ELECTIONS_POLITICS"
	AdcampaigngroupSpecialAdCategoryNone                      AdcampaigngroupSpecialAdCategory = "NONE"
	AdcampaigngroupSpecialAdCategoryOnlineGamblingAndGaming   AdcampaigngroupSpecialAdCategory = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdcampaigngroupSpecialAdCategory) String() string {
	return string(value)
}
