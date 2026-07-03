package enums

type AdgroupinsightsActionAttributionWindowsEnumParam string

const (
	AdgroupinsightsActionAttributionWindowsEnumParamX1dClick                AdgroupinsightsActionAttributionWindowsEnumParam = "1d_click"
	AdgroupinsightsActionAttributionWindowsEnumParamX1dEv                   AdgroupinsightsActionAttributionWindowsEnumParam = "1d_ev"
	AdgroupinsightsActionAttributionWindowsEnumParamX1dView                 AdgroupinsightsActionAttributionWindowsEnumParam = "1d_view"
	AdgroupinsightsActionAttributionWindowsEnumParamX28dClick               AdgroupinsightsActionAttributionWindowsEnumParam = "28d_click"
	AdgroupinsightsActionAttributionWindowsEnumParamX28dView                AdgroupinsightsActionAttributionWindowsEnumParam = "28d_view"
	AdgroupinsightsActionAttributionWindowsEnumParamX28dViewAllConversions  AdgroupinsightsActionAttributionWindowsEnumParam = "28d_view_all_conversions"
	AdgroupinsightsActionAttributionWindowsEnumParamX28dViewFirstConversion AdgroupinsightsActionAttributionWindowsEnumParam = "28d_view_first_conversion"
	AdgroupinsightsActionAttributionWindowsEnumParamX7dClick                AdgroupinsightsActionAttributionWindowsEnumParam = "7d_click"
	AdgroupinsightsActionAttributionWindowsEnumParamX7dView                 AdgroupinsightsActionAttributionWindowsEnumParam = "7d_view"
	AdgroupinsightsActionAttributionWindowsEnumParamX7dViewAllConversions   AdgroupinsightsActionAttributionWindowsEnumParam = "7d_view_all_conversions"
	AdgroupinsightsActionAttributionWindowsEnumParamX7dViewFirstConversion  AdgroupinsightsActionAttributionWindowsEnumParam = "7d_view_first_conversion"
	AdgroupinsightsActionAttributionWindowsEnumParamDda                     AdgroupinsightsActionAttributionWindowsEnumParam = "dda"
	AdgroupinsightsActionAttributionWindowsEnumParamDefault                 AdgroupinsightsActionAttributionWindowsEnumParam = "default"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanClick               AdgroupinsightsActionAttributionWindowsEnumParam = "skan_click"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanClickSecondPostback AdgroupinsightsActionAttributionWindowsEnumParam = "skan_click_second_postback"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanClickThirdPostback  AdgroupinsightsActionAttributionWindowsEnumParam = "skan_click_third_postback"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanView                AdgroupinsightsActionAttributionWindowsEnumParam = "skan_view"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanViewSecondPostback  AdgroupinsightsActionAttributionWindowsEnumParam = "skan_view_second_postback"
	AdgroupinsightsActionAttributionWindowsEnumParamSkanViewThirdPostback   AdgroupinsightsActionAttributionWindowsEnumParam = "skan_view_third_postback"
)

func (value AdgroupinsightsActionAttributionWindowsEnumParam) String() string {
	return string(value)
}
