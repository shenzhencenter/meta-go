package enums

type AdassetcustomizationrulecustomizationspecDevicePlatforms string

const (
	AdassetcustomizationrulecustomizationspecDevicePlatformsConnectedTv AdassetcustomizationrulecustomizationspecDevicePlatforms = "connected_tv"
	AdassetcustomizationrulecustomizationspecDevicePlatformsDesktop     AdassetcustomizationrulecustomizationspecDevicePlatforms = "desktop"
	AdassetcustomizationrulecustomizationspecDevicePlatformsMobile      AdassetcustomizationrulecustomizationspecDevicePlatforms = "mobile"
)

func (value AdassetcustomizationrulecustomizationspecDevicePlatforms) String() string {
	return string(value)
}
