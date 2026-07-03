package enums

type ApplicationAnPlatforms string

const (
	ApplicationAnPlatformsAndroid         ApplicationAnPlatforms = "ANDROID"
	ApplicationAnPlatformsCtv             ApplicationAnPlatforms = "CTV"
	ApplicationAnPlatformsDesktop         ApplicationAnPlatforms = "DESKTOP"
	ApplicationAnPlatformsGalaxy          ApplicationAnPlatforms = "GALAXY"
	ApplicationAnPlatformsInstantArticles ApplicationAnPlatforms = "INSTANT_ARTICLES"
	ApplicationAnPlatformsIos             ApplicationAnPlatforms = "IOS"
	ApplicationAnPlatformsMobileWeb       ApplicationAnPlatforms = "MOBILE_WEB"
	ApplicationAnPlatformsOculus          ApplicationAnPlatforms = "OCULUS"
	ApplicationAnPlatformsUnknown         ApplicationAnPlatforms = "UNKNOWN"
	ApplicationAnPlatformsXiaomi          ApplicationAnPlatforms = "XIAOMI"
)

func (value ApplicationAnPlatforms) String() string {
	return string(value)
}
