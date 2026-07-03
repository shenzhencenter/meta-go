package enums

type AdcampaignMultiOptimizationGoalWeight string

const (
	AdcampaignMultiOptimizationGoalWeightBalanced      AdcampaignMultiOptimizationGoalWeight = "BALANCED"
	AdcampaignMultiOptimizationGoalWeightPreferEvent   AdcampaignMultiOptimizationGoalWeight = "PREFER_EVENT"
	AdcampaignMultiOptimizationGoalWeightPreferInstall AdcampaignMultiOptimizationGoalWeight = "PREFER_INSTALL"
	AdcampaignMultiOptimizationGoalWeightUndefined     AdcampaignMultiOptimizationGoalWeight = "UNDEFINED"
)

func (value AdcampaignMultiOptimizationGoalWeight) String() string {
	return string(value)
}
