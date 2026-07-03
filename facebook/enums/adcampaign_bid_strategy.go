package enums

type AdcampaignBidStrategy string

const (
	AdcampaignBidStrategyCostCap               AdcampaignBidStrategy = "COST_CAP"
	AdcampaignBidStrategyLowestCostWithoutCap  AdcampaignBidStrategy = "LOWEST_COST_WITHOUT_CAP"
	AdcampaignBidStrategyLowestCostWithBidCap  AdcampaignBidStrategy = "LOWEST_COST_WITH_BID_CAP"
	AdcampaignBidStrategyLowestCostWithMinRoas AdcampaignBidStrategy = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value AdcampaignBidStrategy) String() string {
	return string(value)
}
