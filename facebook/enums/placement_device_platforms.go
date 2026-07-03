package enums

type PlacementDevicePlatforms string

const (
	PlacementDevicePlatformsConnectedTv PlacementDevicePlatforms = "connected_tv"
	PlacementDevicePlatformsDesktop     PlacementDevicePlatforms = "desktop"
	PlacementDevicePlatformsMobile      PlacementDevicePlatforms = "mobile"
)

func (value PlacementDevicePlatforms) String() string {
	return string(value)
}
