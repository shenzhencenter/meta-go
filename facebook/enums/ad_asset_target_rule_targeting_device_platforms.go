package enums

type AdassettargetruletargetingDevicePlatforms string

const (
	AdassettargetruletargetingDevicePlatformsConnectedTv AdassettargetruletargetingDevicePlatforms = "connected_tv"
	AdassettargetruletargetingDevicePlatformsDesktop     AdassettargetruletargetingDevicePlatforms = "desktop"
	AdassettargetruletargetingDevicePlatformsMobile      AdassettargetruletargetingDevicePlatforms = "mobile"
)

func (value AdassettargetruletargetingDevicePlatforms) String() string {
	return string(value)
}
