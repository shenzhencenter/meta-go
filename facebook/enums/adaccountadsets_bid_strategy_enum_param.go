package enums

type AdaccountadsetsBidStrategyEnumParam string

const (
	AdaccountadsetsBidStrategyEnumParamCostCap               AdaccountadsetsBidStrategyEnumParam = "COST_CAP"
	AdaccountadsetsBidStrategyEnumParamLowestCostWithoutCap  AdaccountadsetsBidStrategyEnumParam = "LOWEST_COST_WITHOUT_CAP"
	AdaccountadsetsBidStrategyEnumParamLowestCostWithBidCap  AdaccountadsetsBidStrategyEnumParam = "LOWEST_COST_WITH_BID_CAP"
	AdaccountadsetsBidStrategyEnumParamLowestCostWithMinRoas AdaccountadsetsBidStrategyEnumParam = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value AdaccountadsetsBidStrategyEnumParam) String() string {
	return string(value)
}
