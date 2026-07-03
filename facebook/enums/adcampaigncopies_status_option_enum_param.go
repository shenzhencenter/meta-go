package enums

type AdcampaigncopiesStatusOptionEnumParam string

const (
	AdcampaigncopiesStatusOptionEnumParamActive              AdcampaigncopiesStatusOptionEnumParam = "ACTIVE"
	AdcampaigncopiesStatusOptionEnumParamInheritedFromSource AdcampaigncopiesStatusOptionEnumParam = "INHERITED_FROM_SOURCE"
	AdcampaigncopiesStatusOptionEnumParamPaused              AdcampaigncopiesStatusOptionEnumParam = "PAUSED"
)

func (value AdcampaigncopiesStatusOptionEnumParam) String() string {
	return string(value)
}
