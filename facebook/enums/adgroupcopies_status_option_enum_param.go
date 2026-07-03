package enums

type AdgroupcopiesStatusOptionEnumParam string

const (
	AdgroupcopiesStatusOptionEnumParamActive              AdgroupcopiesStatusOptionEnumParam = "ACTIVE"
	AdgroupcopiesStatusOptionEnumParamInheritedFromSource AdgroupcopiesStatusOptionEnumParam = "INHERITED_FROM_SOURCE"
	AdgroupcopiesStatusOptionEnumParamPaused              AdgroupcopiesStatusOptionEnumParam = "PAUSED"
)

func (value AdgroupcopiesStatusOptionEnumParam) String() string {
	return string(value)
}
