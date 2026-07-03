package enums

type AdaccountcampaignsDeleteStrategyEnumParam string

const (
	AdaccountcampaignsDeleteStrategyEnumParamDeleteAny            AdaccountcampaignsDeleteStrategyEnumParam = "DELETE_ANY"
	AdaccountcampaignsDeleteStrategyEnumParamDeleteArchivedBefore AdaccountcampaignsDeleteStrategyEnumParam = "DELETE_ARCHIVED_BEFORE"
	AdaccountcampaignsDeleteStrategyEnumParamDeleteOldest         AdaccountcampaignsDeleteStrategyEnumParam = "DELETE_OLDEST"
)

func (value AdaccountcampaignsDeleteStrategyEnumParam) String() string {
	return string(value)
}
