package enums

type TargetingEffectiveDevicePlatforms string

const (
	TargetingEffectiveDevicePlatformsConnectedTv TargetingEffectiveDevicePlatforms = "connected_tv"
	TargetingEffectiveDevicePlatformsDesktop     TargetingEffectiveDevicePlatforms = "desktop"
	TargetingEffectiveDevicePlatformsMobile      TargetingEffectiveDevicePlatforms = "mobile"
)

func (value TargetingEffectiveDevicePlatforms) String() string {
	return string(value)
}
