package enums

type AdcampaigngroupinsightsActionAttributionWindowsEnumParam string

const (
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX1dClick                AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "1d_click"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX1dEv                   AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "1d_ev"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX1dView                 AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "1d_view"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX28dClick               AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "28d_click"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX28dView                AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "28d_view"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX28dViewAllConversions  AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "28d_view_all_conversions"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX28dViewFirstConversion AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "28d_view_first_conversion"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX7dClick                AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "7d_click"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX7dView                 AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "7d_view"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX7dViewAllConversions   AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "7d_view_all_conversions"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamX7dViewFirstConversion  AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "7d_view_first_conversion"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamDda                     AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "dda"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamDefault                 AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "default"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanClick               AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_click"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanClickSecondPostback AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_click_second_postback"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanClickThirdPostback  AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_click_third_postback"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanView                AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_view"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanViewSecondPostback  AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_view_second_postback"
	AdcampaigngroupinsightsActionAttributionWindowsEnumParamSkanViewThirdPostback   AdcampaigngroupinsightsActionAttributionWindowsEnumParam = "skan_view_third_postback"
)

func (value AdcampaigngroupinsightsActionAttributionWindowsEnumParam) String() string {
	return string(value)
}
