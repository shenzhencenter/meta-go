package enums

type AdaccountadsExecutionOptionsEnumParam string

const (
	AdaccountadsExecutionOptionsEnumParamIncludeRecommendations AdaccountadsExecutionOptionsEnumParam = "include_recommendations"
	AdaccountadsExecutionOptionsEnumParamSynchronousAdReview    AdaccountadsExecutionOptionsEnumParam = "synchronous_ad_review"
	AdaccountadsExecutionOptionsEnumParamValidateOnly           AdaccountadsExecutionOptionsEnumParam = "validate_only"
)

func (value AdaccountadsExecutionOptionsEnumParam) String() string {
	return string(value)
}
