package enums

type ApplicationcodelessEventMappingsPlatformEnumParam string

const (
	ApplicationcodelessEventMappingsPlatformEnumParamAndroid ApplicationcodelessEventMappingsPlatformEnumParam = "ANDROID"
	ApplicationcodelessEventMappingsPlatformEnumParamIos     ApplicationcodelessEventMappingsPlatformEnumParam = "IOS"
)

func (value ApplicationcodelessEventMappingsPlatformEnumParam) String() string {
	return string(value)
}
