package enums

type AdaccountcampaignsEffectiveStatusEnumParam string

const (
	AdaccountcampaignsEffectiveStatusEnumParamActive             AdaccountcampaignsEffectiveStatusEnumParam = "ACTIVE"
	AdaccountcampaignsEffectiveStatusEnumParamAdsetPaused        AdaccountcampaignsEffectiveStatusEnumParam = "ADSET_PAUSED"
	AdaccountcampaignsEffectiveStatusEnumParamArchived           AdaccountcampaignsEffectiveStatusEnumParam = "ARCHIVED"
	AdaccountcampaignsEffectiveStatusEnumParamCampaignPaused     AdaccountcampaignsEffectiveStatusEnumParam = "CAMPAIGN_PAUSED"
	AdaccountcampaignsEffectiveStatusEnumParamDeleted            AdaccountcampaignsEffectiveStatusEnumParam = "DELETED"
	AdaccountcampaignsEffectiveStatusEnumParamDisapproved        AdaccountcampaignsEffectiveStatusEnumParam = "DISAPPROVED"
	AdaccountcampaignsEffectiveStatusEnumParamInProcess          AdaccountcampaignsEffectiveStatusEnumParam = "IN_PROCESS"
	AdaccountcampaignsEffectiveStatusEnumParamPaused             AdaccountcampaignsEffectiveStatusEnumParam = "PAUSED"
	AdaccountcampaignsEffectiveStatusEnumParamPendingBillingInfo AdaccountcampaignsEffectiveStatusEnumParam = "PENDING_BILLING_INFO"
	AdaccountcampaignsEffectiveStatusEnumParamPendingReview      AdaccountcampaignsEffectiveStatusEnumParam = "PENDING_REVIEW"
	AdaccountcampaignsEffectiveStatusEnumParamPreapproved        AdaccountcampaignsEffectiveStatusEnumParam = "PREAPPROVED"
	AdaccountcampaignsEffectiveStatusEnumParamWithIssues         AdaccountcampaignsEffectiveStatusEnumParam = "WITH_ISSUES"
)

func (value AdaccountcampaignsEffectiveStatusEnumParam) String() string {
	return string(value)
}
