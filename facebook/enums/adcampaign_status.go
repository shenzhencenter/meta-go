package enums

type AdcampaignStatus string

const (
	AdcampaignStatusActive   AdcampaignStatus = "ACTIVE"
	AdcampaignStatusArchived AdcampaignStatus = "ARCHIVED"
	AdcampaignStatusDeleted  AdcampaignStatus = "DELETED"
	AdcampaignStatusPaused   AdcampaignStatus = "PAUSED"
)

func (value AdcampaignStatus) String() string {
	return string(value)
}
