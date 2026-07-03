package enums

type AdaccountinsightsActionAttributionWindowsEnumParam string

const (
	AdaccountinsightsActionAttributionWindowsEnumParamX1dClick                AdaccountinsightsActionAttributionWindowsEnumParam = "1d_click"
	AdaccountinsightsActionAttributionWindowsEnumParamX1dEv                   AdaccountinsightsActionAttributionWindowsEnumParam = "1d_ev"
	AdaccountinsightsActionAttributionWindowsEnumParamX1dView                 AdaccountinsightsActionAttributionWindowsEnumParam = "1d_view"
	AdaccountinsightsActionAttributionWindowsEnumParamX28dClick               AdaccountinsightsActionAttributionWindowsEnumParam = "28d_click"
	AdaccountinsightsActionAttributionWindowsEnumParamX28dView                AdaccountinsightsActionAttributionWindowsEnumParam = "28d_view"
	AdaccountinsightsActionAttributionWindowsEnumParamX28dViewAllConversions  AdaccountinsightsActionAttributionWindowsEnumParam = "28d_view_all_conversions"
	AdaccountinsightsActionAttributionWindowsEnumParamX28dViewFirstConversion AdaccountinsightsActionAttributionWindowsEnumParam = "28d_view_first_conversion"
	AdaccountinsightsActionAttributionWindowsEnumParamX7dClick                AdaccountinsightsActionAttributionWindowsEnumParam = "7d_click"
	AdaccountinsightsActionAttributionWindowsEnumParamX7dView                 AdaccountinsightsActionAttributionWindowsEnumParam = "7d_view"
	AdaccountinsightsActionAttributionWindowsEnumParamX7dViewAllConversions   AdaccountinsightsActionAttributionWindowsEnumParam = "7d_view_all_conversions"
	AdaccountinsightsActionAttributionWindowsEnumParamX7dViewFirstConversion  AdaccountinsightsActionAttributionWindowsEnumParam = "7d_view_first_conversion"
	AdaccountinsightsActionAttributionWindowsEnumParamDda                     AdaccountinsightsActionAttributionWindowsEnumParam = "dda"
	AdaccountinsightsActionAttributionWindowsEnumParamDefault                 AdaccountinsightsActionAttributionWindowsEnumParam = "default"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanClick               AdaccountinsightsActionAttributionWindowsEnumParam = "skan_click"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanClickSecondPostback AdaccountinsightsActionAttributionWindowsEnumParam = "skan_click_second_postback"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanClickThirdPostback  AdaccountinsightsActionAttributionWindowsEnumParam = "skan_click_third_postback"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanView                AdaccountinsightsActionAttributionWindowsEnumParam = "skan_view"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanViewSecondPostback  AdaccountinsightsActionAttributionWindowsEnumParam = "skan_view_second_postback"
	AdaccountinsightsActionAttributionWindowsEnumParamSkanViewThirdPostback   AdaccountinsightsActionAttributionWindowsEnumParam = "skan_view_third_postback"
)

func (value AdaccountinsightsActionAttributionWindowsEnumParam) String() string {
	return string(value)
}
