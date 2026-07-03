package enums

type AdcampaigncopiesEffectiveStatusEnumParam string

const (
	AdcampaigncopiesEffectiveStatusEnumParamActive             AdcampaigncopiesEffectiveStatusEnumParam = "ACTIVE"
	AdcampaigncopiesEffectiveStatusEnumParamAdsetPaused        AdcampaigncopiesEffectiveStatusEnumParam = "ADSET_PAUSED"
	AdcampaigncopiesEffectiveStatusEnumParamArchived           AdcampaigncopiesEffectiveStatusEnumParam = "ARCHIVED"
	AdcampaigncopiesEffectiveStatusEnumParamCampaignPaused     AdcampaigncopiesEffectiveStatusEnumParam = "CAMPAIGN_PAUSED"
	AdcampaigncopiesEffectiveStatusEnumParamDeleted            AdcampaigncopiesEffectiveStatusEnumParam = "DELETED"
	AdcampaigncopiesEffectiveStatusEnumParamDisapproved        AdcampaigncopiesEffectiveStatusEnumParam = "DISAPPROVED"
	AdcampaigncopiesEffectiveStatusEnumParamInProcess          AdcampaigncopiesEffectiveStatusEnumParam = "IN_PROCESS"
	AdcampaigncopiesEffectiveStatusEnumParamPaused             AdcampaigncopiesEffectiveStatusEnumParam = "PAUSED"
	AdcampaigncopiesEffectiveStatusEnumParamPendingBillingInfo AdcampaigncopiesEffectiveStatusEnumParam = "PENDING_BILLING_INFO"
	AdcampaigncopiesEffectiveStatusEnumParamPendingReview      AdcampaigncopiesEffectiveStatusEnumParam = "PENDING_REVIEW"
	AdcampaigncopiesEffectiveStatusEnumParamPreapproved        AdcampaigncopiesEffectiveStatusEnumParam = "PREAPPROVED"
	AdcampaigncopiesEffectiveStatusEnumParamWithIssues         AdcampaigncopiesEffectiveStatusEnumParam = "WITH_ISSUES"
)

func (value AdcampaigncopiesEffectiveStatusEnumParam) String() string {
	return string(value)
}
