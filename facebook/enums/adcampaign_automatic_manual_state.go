package enums

type AdcampaignAutomaticManualState string

const (
	AdcampaignAutomaticManualStateAutomatic AdcampaignAutomaticManualState = "AUTOMATIC"
	AdcampaignAutomaticManualStateManual    AdcampaignAutomaticManualState = "MANUAL"
	AdcampaignAutomaticManualStateUnset     AdcampaignAutomaticManualState = "UNSET"
)

func (value AdcampaignAutomaticManualState) String() string {
	return string(value)
}
