package enums

type UserliveVideosSourceEnumParam string

const (
	UserliveVideosSourceEnumParamOwner  UserliveVideosSourceEnumParam = "owner"
	UserliveVideosSourceEnumParamTarget UserliveVideosSourceEnumParam = "target"
)

func (value UserliveVideosSourceEnumParam) String() string {
	return string(value)
}
