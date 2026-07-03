package enums

type ApplicationmobileSdkGkPlatformEnumParam string

const (
	ApplicationmobileSdkGkPlatformEnumParamAndroid ApplicationmobileSdkGkPlatformEnumParam = "ANDROID"
	ApplicationmobileSdkGkPlatformEnumParamIos     ApplicationmobileSdkGkPlatformEnumParam = "IOS"
)

func (value ApplicationmobileSdkGkPlatformEnumParam) String() string {
	return string(value)
}
