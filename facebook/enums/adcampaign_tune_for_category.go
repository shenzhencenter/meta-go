package enums

type AdcampaignTuneForCategory string

const (
	AdcampaignTuneForCategoryCredit                    AdcampaignTuneForCategory = "CREDIT"
	AdcampaignTuneForCategoryEmployment                AdcampaignTuneForCategory = "EMPLOYMENT"
	AdcampaignTuneForCategoryFinancialProductsServices AdcampaignTuneForCategory = "FINANCIAL_PRODUCTS_SERVICES"
	AdcampaignTuneForCategoryHousing                   AdcampaignTuneForCategory = "HOUSING"
	AdcampaignTuneForCategoryIssuesElectionsPolitics   AdcampaignTuneForCategory = "ISSUES_ELECTIONS_POLITICS"
	AdcampaignTuneForCategoryNone                      AdcampaignTuneForCategory = "NONE"
	AdcampaignTuneForCategoryOnlineGamblingAndGaming   AdcampaignTuneForCategory = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdcampaignTuneForCategory) String() string {
	return string(value)
}
