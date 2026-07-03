package enums

type AdgroupExecutionOptions string

const (
	AdgroupExecutionOptionsIncludeRecommendations AdgroupExecutionOptions = "include_recommendations"
	AdgroupExecutionOptionsSynchronousAdReview    AdgroupExecutionOptions = "synchronous_ad_review"
	AdgroupExecutionOptionsValidateOnly           AdgroupExecutionOptions = "validate_only"
)

func (value AdgroupExecutionOptions) String() string {
	return string(value)
}
