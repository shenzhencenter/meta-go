package enums

type AdaccountadsetsTuneForCategoryEnumParam string

const (
	AdaccountadsetsTuneForCategoryEnumParamCredit                    AdaccountadsetsTuneForCategoryEnumParam = "CREDIT"
	AdaccountadsetsTuneForCategoryEnumParamEmployment                AdaccountadsetsTuneForCategoryEnumParam = "EMPLOYMENT"
	AdaccountadsetsTuneForCategoryEnumParamFinancialProductsServices AdaccountadsetsTuneForCategoryEnumParam = "FINANCIAL_PRODUCTS_SERVICES"
	AdaccountadsetsTuneForCategoryEnumParamHousing                   AdaccountadsetsTuneForCategoryEnumParam = "HOUSING"
	AdaccountadsetsTuneForCategoryEnumParamIssuesElectionsPolitics   AdaccountadsetsTuneForCategoryEnumParam = "ISSUES_ELECTIONS_POLITICS"
	AdaccountadsetsTuneForCategoryEnumParamNone                      AdaccountadsetsTuneForCategoryEnumParam = "NONE"
	AdaccountadsetsTuneForCategoryEnumParamOnlineGamblingAndGaming   AdaccountadsetsTuneForCategoryEnumParam = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdaccountadsetsTuneForCategoryEnumParam) String() string {
	return string(value)
}
