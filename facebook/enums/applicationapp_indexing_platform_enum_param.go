package enums

type ApplicationappIndexingPlatformEnumParam string

const (
	ApplicationappIndexingPlatformEnumParamAndroid ApplicationappIndexingPlatformEnumParam = "ANDROID"
	ApplicationappIndexingPlatformEnumParamIos     ApplicationappIndexingPlatformEnumParam = "IOS"
)

func (value ApplicationappIndexingPlatformEnumParam) String() string {
	return string(value)
}
