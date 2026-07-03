package enums

type AdcampaigngroupcopiesStatusOptionEnumParam string

const (
	AdcampaigngroupcopiesStatusOptionEnumParamActive              AdcampaigngroupcopiesStatusOptionEnumParam = "ACTIVE"
	AdcampaigngroupcopiesStatusOptionEnumParamInheritedFromSource AdcampaigngroupcopiesStatusOptionEnumParam = "INHERITED_FROM_SOURCE"
	AdcampaigngroupcopiesStatusOptionEnumParamPaused              AdcampaigngroupcopiesStatusOptionEnumParam = "PAUSED"
)

func (value AdcampaigngroupcopiesStatusOptionEnumParam) String() string {
	return string(value)
}
