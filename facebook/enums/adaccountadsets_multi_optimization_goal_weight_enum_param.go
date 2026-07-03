package enums

type AdaccountadsetsMultiOptimizationGoalWeightEnumParam string

const (
	AdaccountadsetsMultiOptimizationGoalWeightEnumParamBalanced      AdaccountadsetsMultiOptimizationGoalWeightEnumParam = "BALANCED"
	AdaccountadsetsMultiOptimizationGoalWeightEnumParamPreferEvent   AdaccountadsetsMultiOptimizationGoalWeightEnumParam = "PREFER_EVENT"
	AdaccountadsetsMultiOptimizationGoalWeightEnumParamPreferInstall AdaccountadsetsMultiOptimizationGoalWeightEnumParam = "PREFER_INSTALL"
	AdaccountadsetsMultiOptimizationGoalWeightEnumParamUndefined     AdaccountadsetsMultiOptimizationGoalWeightEnumParam = "UNDEFINED"
)

func (value AdaccountadsetsMultiOptimizationGoalWeightEnumParam) String() string {
	return string(value)
}
