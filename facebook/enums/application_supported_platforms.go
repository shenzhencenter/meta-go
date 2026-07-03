package enums

type ApplicationSupportedPlatforms string

const (
	ApplicationSupportedPlatformsAmazon              ApplicationSupportedPlatforms = "AMAZON"
	ApplicationSupportedPlatformsAndroid             ApplicationSupportedPlatforms = "ANDROID"
	ApplicationSupportedPlatformsCanvas              ApplicationSupportedPlatforms = "CANVAS"
	ApplicationSupportedPlatformsGameroom            ApplicationSupportedPlatforms = "GAMEROOM"
	ApplicationSupportedPlatformsInstantGame         ApplicationSupportedPlatforms = "INSTANT_GAME"
	ApplicationSupportedPlatformsIpad                ApplicationSupportedPlatforms = "IPAD"
	ApplicationSupportedPlatformsIphone              ApplicationSupportedPlatforms = "IPHONE"
	ApplicationSupportedPlatformsMobileWeb           ApplicationSupportedPlatforms = "MOBILE_WEB"
	ApplicationSupportedPlatformsOculus              ApplicationSupportedPlatforms = "OCULUS"
	ApplicationSupportedPlatformsSamsung             ApplicationSupportedPlatforms = "SAMSUNG"
	ApplicationSupportedPlatformsSupplementaryImages ApplicationSupportedPlatforms = "SUPPLEMENTARY_IMAGES"
	ApplicationSupportedPlatformsWeb                 ApplicationSupportedPlatforms = "WEB"
	ApplicationSupportedPlatformsWindows             ApplicationSupportedPlatforms = "WINDOWS"
	ApplicationSupportedPlatformsXiaomi              ApplicationSupportedPlatforms = "XIAOMI"
)

func (value ApplicationSupportedPlatforms) String() string {
	return string(value)
}
