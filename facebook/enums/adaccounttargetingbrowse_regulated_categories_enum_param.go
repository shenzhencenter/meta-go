package enums

type AdaccounttargetingbrowseRegulatedCategoriesEnumParam string

const (
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamCredit                    AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "CREDIT"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamEmployment                AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "EMPLOYMENT"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamFinancialProductsServices AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "FINANCIAL_PRODUCTS_SERVICES"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamHousing                   AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "HOUSING"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamIssuesElectionsPolitics   AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "ISSUES_ELECTIONS_POLITICS"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamNone                      AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "NONE"
	AdaccounttargetingbrowseRegulatedCategoriesEnumParamOnlineGamblingAndGaming   AdaccounttargetingbrowseRegulatedCategoriesEnumParam = "ONLINE_GAMBLING_AND_GAMING"
)

func (value AdaccounttargetingbrowseRegulatedCategoriesEnumParam) String() string {
	return string(value)
}
