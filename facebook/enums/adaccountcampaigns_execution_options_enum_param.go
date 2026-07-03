package enums

type AdaccountcampaignsExecutionOptionsEnumParam string

const (
	AdaccountcampaignsExecutionOptionsEnumParamIncludeRecommendations AdaccountcampaignsExecutionOptionsEnumParam = "include_recommendations"
	AdaccountcampaignsExecutionOptionsEnumParamValidateOnly           AdaccountcampaignsExecutionOptionsEnumParam = "validate_only"
)

func (value AdaccountcampaignsExecutionOptionsEnumParam) String() string {
	return string(value)
}
