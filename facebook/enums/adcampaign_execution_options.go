package enums

type AdcampaignExecutionOptions string

const (
	AdcampaignExecutionOptionsIncludeRecommendations AdcampaignExecutionOptions = "include_recommendations"
	AdcampaignExecutionOptionsValidateOnly           AdcampaignExecutionOptions = "validate_only"
)

func (value AdcampaignExecutionOptions) String() string {
	return string(value)
}
