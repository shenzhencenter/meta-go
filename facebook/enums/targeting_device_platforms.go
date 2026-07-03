package enums

type TargetingDevicePlatforms string

const (
	TargetingDevicePlatformsConnectedTv TargetingDevicePlatforms = "connected_tv"
	TargetingDevicePlatformsDesktop     TargetingDevicePlatforms = "desktop"
	TargetingDevicePlatformsMobile      TargetingDevicePlatforms = "mobile"
)

func (value TargetingDevicePlatforms) String() string {
	return string(value)
}
