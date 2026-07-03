package enums

type AdEffectiveStatus string

const (
	AdEffectiveStatusActive             AdEffectiveStatus = "ACTIVE"
	AdEffectiveStatusAdsetPaused        AdEffectiveStatus = "ADSET_PAUSED"
	AdEffectiveStatusArchived           AdEffectiveStatus = "ARCHIVED"
	AdEffectiveStatusCampaignPaused     AdEffectiveStatus = "CAMPAIGN_PAUSED"
	AdEffectiveStatusDeleted            AdEffectiveStatus = "DELETED"
	AdEffectiveStatusDisapproved        AdEffectiveStatus = "DISAPPROVED"
	AdEffectiveStatusInProcess          AdEffectiveStatus = "IN_PROCESS"
	AdEffectiveStatusPaused             AdEffectiveStatus = "PAUSED"
	AdEffectiveStatusPendingBillingInfo AdEffectiveStatus = "PENDING_BILLING_INFO"
	AdEffectiveStatusPendingReview      AdEffectiveStatus = "PENDING_REVIEW"
	AdEffectiveStatusPreapproved        AdEffectiveStatus = "PREAPPROVED"
	AdEffectiveStatusWithIssues         AdEffectiveStatus = "WITH_ISSUES"
)

func (value AdEffectiveStatus) String() string {
	return string(value)
}
