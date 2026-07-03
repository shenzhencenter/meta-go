package enums

type AdaccounttargetingsuggestionsModeEnumParam string

const (
	AdaccounttargetingsuggestionsModeEnumParamBestPerforming AdaccounttargetingsuggestionsModeEnumParam = "best_performing"
	AdaccounttargetingsuggestionsModeEnumParamRecentlyUsed   AdaccounttargetingsuggestionsModeEnumParam = "recently_used"
	AdaccounttargetingsuggestionsModeEnumParamRelated        AdaccounttargetingsuggestionsModeEnumParam = "related"
	AdaccounttargetingsuggestionsModeEnumParamSuggestions    AdaccounttargetingsuggestionsModeEnumParam = "suggestions"
)

func (value AdaccounttargetingsuggestionsModeEnumParam) String() string {
	return string(value)
}
