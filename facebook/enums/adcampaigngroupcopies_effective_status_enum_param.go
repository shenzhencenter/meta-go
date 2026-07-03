package enums

type AdcampaigngroupcopiesEffectiveStatusEnumParam string

const (
	AdcampaigngroupcopiesEffectiveStatusEnumParamActive             AdcampaigngroupcopiesEffectiveStatusEnumParam = "ACTIVE"
	AdcampaigngroupcopiesEffectiveStatusEnumParamAdsetPaused        AdcampaigngroupcopiesEffectiveStatusEnumParam = "ADSET_PAUSED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamArchived           AdcampaigngroupcopiesEffectiveStatusEnumParam = "ARCHIVED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamCampaignPaused     AdcampaigngroupcopiesEffectiveStatusEnumParam = "CAMPAIGN_PAUSED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamDeleted            AdcampaigngroupcopiesEffectiveStatusEnumParam = "DELETED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamDisapproved        AdcampaigngroupcopiesEffectiveStatusEnumParam = "DISAPPROVED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamInProcess          AdcampaigngroupcopiesEffectiveStatusEnumParam = "IN_PROCESS"
	AdcampaigngroupcopiesEffectiveStatusEnumParamPaused             AdcampaigngroupcopiesEffectiveStatusEnumParam = "PAUSED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamPendingBillingInfo AdcampaigngroupcopiesEffectiveStatusEnumParam = "PENDING_BILLING_INFO"
	AdcampaigngroupcopiesEffectiveStatusEnumParamPendingReview      AdcampaigngroupcopiesEffectiveStatusEnumParam = "PENDING_REVIEW"
	AdcampaigngroupcopiesEffectiveStatusEnumParamPreapproved        AdcampaigngroupcopiesEffectiveStatusEnumParam = "PREAPPROVED"
	AdcampaigngroupcopiesEffectiveStatusEnumParamWithIssues         AdcampaigngroupcopiesEffectiveStatusEnumParam = "WITH_ISSUES"
)

func (value AdcampaigngroupcopiesEffectiveStatusEnumParam) String() string {
	return string(value)
}
