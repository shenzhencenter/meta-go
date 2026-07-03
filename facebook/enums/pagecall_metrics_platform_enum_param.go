package enums

type PagecallMetricsPlatformEnumParam string

const (
	PagecallMetricsPlatformEnumParamMessenger PagecallMetricsPlatformEnumParam = "MESSENGER"
)

func (value PagecallMetricsPlatformEnumParam) String() string {
	return string(value)
}
