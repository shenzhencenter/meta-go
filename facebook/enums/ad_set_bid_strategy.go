package enums

type AdsetBidStrategy string

const (
	AdsetBidStrategyCostCap               AdsetBidStrategy = "COST_CAP"
	AdsetBidStrategyLowestCostWithoutCap  AdsetBidStrategy = "LOWEST_COST_WITHOUT_CAP"
	AdsetBidStrategyLowestCostWithBidCap  AdsetBidStrategy = "LOWEST_COST_WITH_BID_CAP"
	AdsetBidStrategyLowestCostWithMinRoas AdsetBidStrategy = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value AdsetBidStrategy) String() string {
	return string(value)
}
