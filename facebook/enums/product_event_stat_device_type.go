package enums

type ProducteventstatDeviceType string

const (
	ProducteventstatDeviceTypeDesktop             ProducteventstatDeviceType = "desktop"
	ProducteventstatDeviceTypeMobileAndroidPhone  ProducteventstatDeviceType = "mobile_android_phone"
	ProducteventstatDeviceTypeMobileAndroidTablet ProducteventstatDeviceType = "mobile_android_tablet"
	ProducteventstatDeviceTypeMobileIpad          ProducteventstatDeviceType = "mobile_ipad"
	ProducteventstatDeviceTypeMobileIphone        ProducteventstatDeviceType = "mobile_iphone"
	ProducteventstatDeviceTypeMobileIpod          ProducteventstatDeviceType = "mobile_ipod"
	ProducteventstatDeviceTypeMobilePhone         ProducteventstatDeviceType = "mobile_phone"
	ProducteventstatDeviceTypeMobileTablet        ProducteventstatDeviceType = "mobile_tablet"
	ProducteventstatDeviceTypeMobileWindowsPhone  ProducteventstatDeviceType = "mobile_windows_phone"
	ProducteventstatDeviceTypeUnknown             ProducteventstatDeviceType = "unknown"
)

func (value ProducteventstatDeviceType) String() string {
	return string(value)
}
