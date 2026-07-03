package enums

type PagecallsPlatformEnumParam string

const (
	PagecallsPlatformEnumParamInstagram PagecallsPlatformEnumParam = "INSTAGRAM"
	PagecallsPlatformEnumParamMessenger PagecallsPlatformEnumParam = "MESSENGER"
)

func (value PagecallsPlatformEnumParam) String() string {
	return string(value)
}
