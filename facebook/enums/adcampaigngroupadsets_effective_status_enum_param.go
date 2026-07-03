package enums

type AdcampaigngroupadsetsEffectiveStatusEnumParam string

const (
	AdcampaigngroupadsetsEffectiveStatusEnumParamActive             AdcampaigngroupadsetsEffectiveStatusEnumParam = "ACTIVE"
	AdcampaigngroupadsetsEffectiveStatusEnumParamAdsetPaused        AdcampaigngroupadsetsEffectiveStatusEnumParam = "ADSET_PAUSED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamArchived           AdcampaigngroupadsetsEffectiveStatusEnumParam = "ARCHIVED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamCampaignPaused     AdcampaigngroupadsetsEffectiveStatusEnumParam = "CAMPAIGN_PAUSED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamDeleted            AdcampaigngroupadsetsEffectiveStatusEnumParam = "DELETED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamDisapproved        AdcampaigngroupadsetsEffectiveStatusEnumParam = "DISAPPROVED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamInProcess          AdcampaigngroupadsetsEffectiveStatusEnumParam = "IN_PROCESS"
	AdcampaigngroupadsetsEffectiveStatusEnumParamPaused             AdcampaigngroupadsetsEffectiveStatusEnumParam = "PAUSED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamPendingBillingInfo AdcampaigngroupadsetsEffectiveStatusEnumParam = "PENDING_BILLING_INFO"
	AdcampaigngroupadsetsEffectiveStatusEnumParamPendingReview      AdcampaigngroupadsetsEffectiveStatusEnumParam = "PENDING_REVIEW"
	AdcampaigngroupadsetsEffectiveStatusEnumParamPreapproved        AdcampaigngroupadsetsEffectiveStatusEnumParam = "PREAPPROVED"
	AdcampaigngroupadsetsEffectiveStatusEnumParamWithIssues         AdcampaigngroupadsetsEffectiveStatusEnumParam = "WITH_ISSUES"
)

func (value AdcampaigngroupadsetsEffectiveStatusEnumParam) String() string {
	return string(value)
}
