package enums

type AdaccountadsetsExecutionOptionsEnumParam string

const (
	AdaccountadsetsExecutionOptionsEnumParamIncludeRecommendations AdaccountadsetsExecutionOptionsEnumParam = "include_recommendations"
	AdaccountadsetsExecutionOptionsEnumParamValidateOnly           AdaccountadsetsExecutionOptionsEnumParam = "validate_only"
)

func (value AdaccountadsetsExecutionOptionsEnumParam) String() string {
	return string(value)
}
