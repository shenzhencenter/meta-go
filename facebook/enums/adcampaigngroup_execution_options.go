package enums

type AdcampaigngroupExecutionOptions string

const (
	AdcampaigngroupExecutionOptionsIncludeRecommendations AdcampaigngroupExecutionOptions = "include_recommendations"
	AdcampaigngroupExecutionOptionsValidateOnly           AdcampaigngroupExecutionOptions = "validate_only"
)

func (value AdcampaigngroupExecutionOptions) String() string {
	return string(value)
}
