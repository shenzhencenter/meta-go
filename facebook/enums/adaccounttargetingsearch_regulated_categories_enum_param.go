package enums

type AdaccounttargetingsearchRegulatedCategoriesEnumParam string

const (
	AdaccounttargetingsearchRegulatedCategoriesEnumParamCredit                    AdaccounttargetingsearchRegulatedCategoriesEnumParam = "CREDIT"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamEmployment                AdaccounttargetingsearchRegulatedCategoriesEnumParam = "EMPLOYMENT"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamFinancialProductsServices AdaccounttargetingsearchRegulatedCategoriesEnumParam = "FINANCIAL_PRODUCTS_SERVICES"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamHousing                   AdaccounttargetingsearchRegulatedCategoriesEnumParam = "HOUSING"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamIssuesElectionsPolitics   AdaccounttargetingsearchRegulatedCategoriesEnumParam = "ISSUES_ELECTIONS_POLITICS"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamNone                      AdaccounttargetingsearchRegulatedCategoriesEnumParam = "NONE"
	AdaccounttargetingsearchRegulatedCategoriesEnumParamOnlineGamblingAndGaming   AdaccounttargetingsearchRegulatedCategoriesEnumParam = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdaccounttargetingsearchRegulatedCategoriesEnumParam) String() string {
	return string(value)
}
