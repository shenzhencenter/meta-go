package enums

type PlacementEffectiveDevicePlatforms string

const (
	PlacementEffectiveDevicePlatformsConnectedTv PlacementEffectiveDevicePlatforms = "connected_tv"
	PlacementEffectiveDevicePlatformsDesktop     PlacementEffectiveDevicePlatforms = "desktop"
	PlacementEffectiveDevicePlatformsMobile      PlacementEffectiveDevicePlatforms = "mobile"
)

func (value PlacementEffectiveDevicePlatforms) String() string {
	return string(value)
}
