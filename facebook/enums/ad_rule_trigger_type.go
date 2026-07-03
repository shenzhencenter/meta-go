package enums

type AdruletriggerType string

const (
	AdruletriggerTypeDeliveryInsightsChange AdruletriggerType = "DELIVERY_INSIGHTS_CHANGE"
	AdruletriggerTypeMetadataCreation       AdruletriggerType = "METADATA_CREATION"
	AdruletriggerTypeMetadataUpdate         AdruletriggerType = "METADATA_UPDATE"
	AdruletriggerTypeStatsChange            AdruletriggerType = "STATS_CHANGE"
	AdruletriggerTypeStatsMilestone         AdruletriggerType = "STATS_MILESTONE"
)

func (value AdruletriggerType) String() string {
	return string(value)
}
