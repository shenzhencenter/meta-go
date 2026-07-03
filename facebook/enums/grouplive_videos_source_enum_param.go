package enums

type GroupliveVideosSourceEnumParam string

const (
	GroupliveVideosSourceEnumParamOwner  GroupliveVideosSourceEnumParam = "owner"
	GroupliveVideosSourceEnumParamTarget GroupliveVideosSourceEnumParam = "target"
)

func (value GroupliveVideosSourceEnumParam) String() string {
	return string(value)
}
