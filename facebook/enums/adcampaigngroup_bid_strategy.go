package enums

type AdcampaigngroupBidStrategy string

const (
	AdcampaigngroupBidStrategyCostCap               AdcampaigngroupBidStrategy = "COST_CAP"
	AdcampaigngroupBidStrategyLowestCostWithoutCap  AdcampaigngroupBidStrategy = "LOWEST_COST_WITHOUT_CAP"
	AdcampaigngroupBidStrategyLowestCostWithBidCap  AdcampaigngroupBidStrategy = "LOWEST_COST_WITH_BID_CAP"
	AdcampaigngroupBidStrategyLowestCostWithMinRoas AdcampaigngroupBidStrategy = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value AdcampaigngroupBidStrategy) String() string {
	return string(value)
}
