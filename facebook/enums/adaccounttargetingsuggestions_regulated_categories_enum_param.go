package enums

type AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam string

const (
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamCredit                    AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "CREDIT"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamEmployment                AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "EMPLOYMENT"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamFinancialProductsServices AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "FINANCIAL_PRODUCTS_SERVICES"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamHousing                   AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "HOUSING"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamIssuesElectionsPolitics   AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "ISSUES_ELECTIONS_POLITICS"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamNone                      AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "NONE"
	AdaccounttargetingsuggestionsRegulatedCategoriesEnumParamOnlineGamblingAndGaming   AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam) String() string {
	return string(value)
}
