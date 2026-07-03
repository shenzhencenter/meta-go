package enums

type ApplicationappPushDeviceTokenPlatformEnumParam string

const (
	ApplicationappPushDeviceTokenPlatformEnumParamAndroid ApplicationappPushDeviceTokenPlatformEnumParam = "ANDROID"
	ApplicationappPushDeviceTokenPlatformEnumParamIos     ApplicationappPushDeviceTokenPlatformEnumParam = "IOS"
	ApplicationappPushDeviceTokenPlatformEnumParamUnknown ApplicationappPushDeviceTokenPlatformEnumParam = "UNKNOWN"
)

func (value ApplicationappPushDeviceTokenPlatformEnumParam) String() string {
	return string(value)
}
