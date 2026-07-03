package enums

type AdruleexecutionspecExecutionType string

const (
	AdruleexecutionspecExecutionTypeAddInterestRelaxation         AdruleexecutionspecExecutionType = "ADD_INTEREST_RELAXATION"
	AdruleexecutionspecExecutionTypeAddQuestionnaireInterests     AdruleexecutionspecExecutionType = "ADD_QUESTIONNAIRE_INTERESTS"
	AdruleexecutionspecExecutionTypeAdRecommendationApply         AdruleexecutionspecExecutionType = "AD_RECOMMENDATION_APPLY"
	AdruleexecutionspecExecutionTypeAudienceConsolidation         AdruleexecutionspecExecutionType = "AUDIENCE_CONSOLIDATION"
	AdruleexecutionspecExecutionTypeAudienceConsolidationAskFirst AdruleexecutionspecExecutionType = "AUDIENCE_CONSOLIDATION_ASK_FIRST"
	AdruleexecutionspecExecutionTypeChangeBid                     AdruleexecutionspecExecutionType = "CHANGE_BID"
	AdruleexecutionspecExecutionTypeChangeBudget                  AdruleexecutionspecExecutionType = "CHANGE_BUDGET"
	AdruleexecutionspecExecutionTypeChangeCampaignBudget          AdruleexecutionspecExecutionType = "CHANGE_CAMPAIGN_BUDGET"
	AdruleexecutionspecExecutionTypeDco                           AdruleexecutionspecExecutionType = "DCO"
	AdruleexecutionspecExecutionTypeIncreaseRadius                AdruleexecutionspecExecutionType = "INCREASE_RADIUS"
	AdruleexecutionspecExecutionTypeNotification                  AdruleexecutionspecExecutionType = "NOTIFICATION"
	AdruleexecutionspecExecutionTypePause                         AdruleexecutionspecExecutionType = "PAUSE"
	AdruleexecutionspecExecutionTypePingEndpoint                  AdruleexecutionspecExecutionType = "PING_ENDPOINT"
	AdruleexecutionspecExecutionTypeRebalanceBudget               AdruleexecutionspecExecutionType = "REBALANCE_BUDGET"
	AdruleexecutionspecExecutionTypeRotate                        AdruleexecutionspecExecutionType = "ROTATE"
	AdruleexecutionspecExecutionTypeUnpause                       AdruleexecutionspecExecutionType = "UNPAUSE"
	AdruleexecutionspecExecutionTypeUpdateCreative                AdruleexecutionspecExecutionType = "UPDATE_CREATIVE"
	AdruleexecutionspecExecutionTypeUpdateLaxBudget               AdruleexecutionspecExecutionType = "UPDATE_LAX_BUDGET"
	AdruleexecutionspecExecutionTypeUpdateLaxDuration             AdruleexecutionspecExecutionType = "UPDATE_LAX_DURATION"
)

func (value AdruleexecutionspecExecutionType) String() string {
	return string(value)
}
