package enums

type AdaccountreachfrequencypredictionsBuyingTypeEnumParam string

const (
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamAuction              AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "AUCTION"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamDeprecatedReachBlock AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "DEPRECATED_REACH_BLOCK"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamFixedCpm             AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "FIXED_CPM"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamMixed                AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "MIXED"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamReachblock           AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "REACHBLOCK"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamResearchPoll         AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "RESEARCH_POLL"
	AdaccountreachfrequencypredictionsBuyingTypeEnumParamReserved             AdaccountreachfrequencypredictionsBuyingTypeEnumParam = "RESERVED"
)

func (value AdaccountreachfrequencypredictionsBuyingTypeEnumParam) String() string {
	return string(value)
}
