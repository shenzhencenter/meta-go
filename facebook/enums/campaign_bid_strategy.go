package enums

type CampaignBidStrategy string

const (
	CampaignBidStrategyCostCap               CampaignBidStrategy = "COST_CAP"
	CampaignBidStrategyLowestCostWithoutCap  CampaignBidStrategy = "LOWEST_COST_WITHOUT_CAP"
	CampaignBidStrategyLowestCostWithBidCap  CampaignBidStrategy = "LOWEST_COST_WITH_BID_CAP"
	CampaignBidStrategyLowestCostWithMinRoas CampaignBidStrategy = "LOWEST_COST_WITH_MIN_ROAS"
)

func (value CampaignBidStrategy) String() string {
	return string(value)
}
