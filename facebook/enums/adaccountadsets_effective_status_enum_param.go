package enums

type AdaccountadsetsEffectiveStatusEnumParam string

const (
	AdaccountadsetsEffectiveStatusEnumParamActive             AdaccountadsetsEffectiveStatusEnumParam = "ACTIVE"
	AdaccountadsetsEffectiveStatusEnumParamAdsetPaused        AdaccountadsetsEffectiveStatusEnumParam = "ADSET_PAUSED"
	AdaccountadsetsEffectiveStatusEnumParamArchived           AdaccountadsetsEffectiveStatusEnumParam = "ARCHIVED"
	AdaccountadsetsEffectiveStatusEnumParamCampaignPaused     AdaccountadsetsEffectiveStatusEnumParam = "CAMPAIGN_PAUSED"
	AdaccountadsetsEffectiveStatusEnumParamDeleted            AdaccountadsetsEffectiveStatusEnumParam = "DELETED"
	AdaccountadsetsEffectiveStatusEnumParamDisapproved        AdaccountadsetsEffectiveStatusEnumParam = "DISAPPROVED"
	AdaccountadsetsEffectiveStatusEnumParamInProcess          AdaccountadsetsEffectiveStatusEnumParam = "IN_PROCESS"
	AdaccountadsetsEffectiveStatusEnumParamPaused             AdaccountadsetsEffectiveStatusEnumParam = "PAUSED"
	AdaccountadsetsEffectiveStatusEnumParamPendingBillingInfo AdaccountadsetsEffectiveStatusEnumParam = "PENDING_BILLING_INFO"
	AdaccountadsetsEffectiveStatusEnumParamPendingReview      AdaccountadsetsEffectiveStatusEnumParam = "PENDING_REVIEW"
	AdaccountadsetsEffectiveStatusEnumParamPreapproved        AdaccountadsetsEffectiveStatusEnumParam = "PREAPPROVED"
	AdaccountadsetsEffectiveStatusEnumParamWithIssues         AdaccountadsetsEffectiveStatusEnumParam = "WITH_ISSUES"
)

func (value AdaccountadsetsEffectiveStatusEnumParam) String() string {
	return string(value)
}
