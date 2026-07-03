package enums

type AdaccountcampaignsBidStrategyEnumParam string

const (
	AdaccountcampaignsBidStrategyEnumParamCostCap               AdaccountcampaignsBidStrategyEnumParam = "COST_CAP"
	AdaccountcampaignsBidStrategyEnumParamLowestCostWithoutCap  AdaccountcampaignsBidStrategyEnumParam = "LOWEST_COST_WITHOUT_CAP"
	AdaccountcampaignsBidStrategyEnumParamLowestCostWithBidCap  AdaccountcampaignsBidStrategyEnumParam = "LOWEST_COST_WITH_BID_CAP"
	AdaccountcampaignsBidStrategyEnumParamLowestCostWithMinRoas AdaccountcampaignsBidStrategyEnumParam = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value AdaccountcampaignsBidStrategyEnumParam) String() string {
	return string(value)
}
