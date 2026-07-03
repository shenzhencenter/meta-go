package enums

type AdspixelDataUseSetting string

const (
	AdspixelDataUseSettingAdvertisingAndAnalytics AdspixelDataUseSetting = "ADVERTISING_AND_ANALYTICS"
	AdspixelDataUseSettingAnalyticsOnly           AdspixelDataUseSetting = "ANALYTICS_ONLY"
	AdspixelDataUseSettingEmpty                   AdspixelDataUseSetting = "EMPTY"
)

func (value AdspixelDataUseSetting) String() string {
	return string(value)
}
