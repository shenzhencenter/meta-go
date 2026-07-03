package enums

type PageliveVideosSourceEnumParam string

const (
	PageliveVideosSourceEnumParamOwner  PageliveVideosSourceEnumParam = "owner"
	PageliveVideosSourceEnumParamTarget PageliveVideosSourceEnumParam = "target"
)

func (value PageliveVideosSourceEnumParam) String() string {
	return string(value)
}
