package enums

type PagethreadsPlatformEnumParam string

const (
	PagethreadsPlatformEnumParamInstagram PagethreadsPlatformEnumParam = "INSTAGRAM"
	PagethreadsPlatformEnumParamMessenger PagethreadsPlatformEnumParam = "MESSENGER"
)

func (value PagethreadsPlatformEnumParam) String() string {
	return string(value)
}
