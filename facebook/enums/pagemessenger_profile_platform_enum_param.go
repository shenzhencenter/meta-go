package enums

type PagemessengerProfilePlatformEnumParam string

const (
	PagemessengerProfilePlatformEnumParamInstagram PagemessengerProfilePlatformEnumParam = "INSTAGRAM"
	PagemessengerProfilePlatformEnumParamMessenger PagemessengerProfilePlatformEnumParam = "MESSENGER"
)

func (value PagemessengerProfilePlatformEnumParam) String() string {
	return string(value)
}
