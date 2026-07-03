package enums

type AdcampaigninsightsActionAttributionWindowsEnumParam string

const (
	AdcampaigninsightsActionAttributionWindowsEnumParamX1dClick                AdcampaigninsightsActionAttributionWindowsEnumParam = "1d_click"
	AdcampaigninsightsActionAttributionWindowsEnumParamX1dEv                   AdcampaigninsightsActionAttributionWindowsEnumParam = "1d_ev"
	AdcampaigninsightsActionAttributionWindowsEnumParamX1dView                 AdcampaigninsightsActionAttributionWindowsEnumParam = "1d_view"
	AdcampaigninsightsActionAttributionWindowsEnumParamX28dClick               AdcampaigninsightsActionAttributionWindowsEnumParam = "28d_click"
	AdcampaigninsightsActionAttributionWindowsEnumParamX28dView                AdcampaigninsightsActionAttributionWindowsEnumParam = "28d_view"
	AdcampaigninsightsActionAttributionWindowsEnumParamX28dViewAllConversions  AdcampaigninsightsActionAttributionWindowsEnumParam = "28d_view_all_conversions"
	AdcampaigninsightsActionAttributionWindowsEnumParamX28dViewFirstConversion AdcampaigninsightsActionAttributionWindowsEnumParam = "28d_view_first_conversion"
	AdcampaigninsightsActionAttributionWindowsEnumParamX7dClick                AdcampaigninsightsActionAttributionWindowsEnumParam = "7d_click"
	AdcampaigninsightsActionAttributionWindowsEnumParamX7dView                 AdcampaigninsightsActionAttributionWindowsEnumParam = "7d_view"
	AdcampaigninsightsActionAttributionWindowsEnumParamX7dViewAllConversions   AdcampaigninsightsActionAttributionWindowsEnumParam = "7d_view_all_conversions"
	AdcampaigninsightsActionAttributionWindowsEnumParamX7dViewFirstConversion  AdcampaigninsightsActionAttributionWindowsEnumParam = "7d_view_first_conversion"
	AdcampaigninsightsActionAttributionWindowsEnumParamDda                     AdcampaigninsightsActionAttributionWindowsEnumParam = "dda"
	AdcampaigninsightsActionAttributionWindowsEnumParamDefault                 AdcampaigninsightsActionAttributionWindowsEnumParam = "default"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanClick               AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_click"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanClickSecondPostback AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_click_second_postback"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanClickThirdPostback  AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_click_third_postback"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanView                AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_view"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanViewSecondPostback  AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_view_second_postback"
	AdcampaigninsightsActionAttributionWindowsEnumParamSkanViewThirdPostback   AdcampaigninsightsActionAttributionWindowsEnumParam = "skan_view_third_postback"
)

func (value AdcampaigninsightsActionAttributionWindowsEnumParam) String() string {
	return string(value)
}
