package enums

type AdcampaigngroupStatus string

const (
	AdcampaigngroupStatusActive   AdcampaigngroupStatus = "ACTIVE"
	AdcampaigngroupStatusArchived AdcampaigngroupStatus = "ARCHIVED"
	AdcampaigngroupStatusDeleted  AdcampaigngroupStatus = "DELETED"
	AdcampaigngroupStatusPaused   AdcampaigngroupStatus = "PAUSED"
)

func (value AdcampaigngroupStatus) String() string {
	return string(value)
}
