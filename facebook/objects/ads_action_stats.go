package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsActionStats struct {
	X1dClick                            *string  `json:"1d_click,omitempty"`
	X1dClickAllConversions              *string  `json:"1d_click_all_conversions,omitempty"`
	X1dClickFirstConversion             *string  `json:"1d_click_first_conversion,omitempty"`
	X1dEv                               *string  `json:"1d_ev,omitempty"`
	X1dEvAllConversions                 *string  `json:"1d_ev_all_conversions,omitempty"`
	X1dEvFirstConversion                *string  `json:"1d_ev_first_conversion,omitempty"`
	X1dPassback                         *string  `json:"1d_passback,omitempty"`
	X1dView                             *string  `json:"1d_view,omitempty"`
	X1dViewAllConversions               *string  `json:"1d_view_all_conversions,omitempty"`
	X1dViewFirstConversion              *string  `json:"1d_view_first_conversion,omitempty"`
	X28dClick                           *string  `json:"28d_click,omitempty"`
	X28dClickAllConversions             *string  `json:"28d_click_all_conversions,omitempty"`
	X28dClickFirstConversion            *string  `json:"28d_click_first_conversion,omitempty"`
	X28dPassback                        *string  `json:"28d_passback,omitempty"`
	X28dView                            *string  `json:"28d_view,omitempty"`
	X28dViewAllConversions              *string  `json:"28d_view_all_conversions,omitempty"`
	X28dViewFirstConversion             *string  `json:"28d_view_first_conversion,omitempty"`
	X7dClick                            *string  `json:"7d_click,omitempty"`
	X7dClickAllConversions              *string  `json:"7d_click_all_conversions,omitempty"`
	X7dClickFirstConversion             *string  `json:"7d_click_first_conversion,omitempty"`
	X7dPassback                         *string  `json:"7d_passback,omitempty"`
	X7dView                             *string  `json:"7d_view,omitempty"`
	X7dViewAllConversions               *string  `json:"7d_view_all_conversions,omitempty"`
	X7dViewFirstConversion              *string  `json:"7d_view_first_conversion,omitempty"`
	ActionBrand                         *string  `json:"action_brand,omitempty"`
	ActionCanvasComponentID             *core.ID `json:"action_canvas_component_id,omitempty"`
	ActionCanvasComponentName           *string  `json:"action_canvas_component_name,omitempty"`
	ActionCarouselCardID                *core.ID `json:"action_carousel_card_id,omitempty"`
	ActionCarouselCardName              *string  `json:"action_carousel_card_name,omitempty"`
	ActionCategory                      *string  `json:"action_category,omitempty"`
	ActionConvertedProductID            *core.ID `json:"action_converted_product_id,omitempty"`
	ActionDestination                   *string  `json:"action_destination,omitempty"`
	ActionDevice                        *string  `json:"action_device,omitempty"`
	ActionEventChannel                  *string  `json:"action_event_channel,omitempty"`
	ActionLinkClickDestination          *string  `json:"action_link_click_destination,omitempty"`
	ActionLocationCode                  *string  `json:"action_location_code,omitempty"`
	ActionReaction                      *string  `json:"action_reaction,omitempty"`
	ActionTargetID                      *core.ID `json:"action_target_id,omitempty"`
	ActionType                          *string  `json:"action_type,omitempty"`
	ActionVideoAssetID                  *core.ID `json:"action_video_asset_id,omitempty"`
	ActionVideoSound                    *string  `json:"action_video_sound,omitempty"`
	ActionVideoType                     *string  `json:"action_video_type,omitempty"`
	Custom                              *string  `json:"custom,omitempty"`
	Dda                                 *string  `json:"dda,omitempty"`
	Incrementality                      *string  `json:"incrementality,omitempty"`
	IncrementalityAllConversions        *string  `json:"incrementality_all_conversions,omitempty"`
	IncrementalityFirstConversion       *string  `json:"incrementality_first_conversion,omitempty"`
	Inline                              *string  `json:"inline,omitempty"`
	InteractiveComponentStickerID       *core.ID `json:"interactive_component_sticker_id,omitempty"`
	InteractiveComponentStickerResponse *string  `json:"interactive_component_sticker_response,omitempty"`
	PromotedProductSetResult            *string  `json:"promoted_product_set_result,omitempty"`
	SkanClick                           *string  `json:"skan_click,omitempty"`
	SkanClickSecondPostback             *string  `json:"skan_click_second_postback,omitempty"`
	SkanClickThirdPostback              *string  `json:"skan_click_third_postback,omitempty"`
	SkanView                            *string  `json:"skan_view,omitempty"`
	SkanViewSecondPostback              *string  `json:"skan_view_second_postback,omitempty"`
	SkanViewThirdPostback               *string  `json:"skan_view_third_postback,omitempty"`
	Value                               *string  `json:"value,omitempty"`
}

var AdsActionStatsFields = struct {
	X1dClick                            string
	X1dClickAllConversions              string
	X1dClickFirstConversion             string
	X1dEv                               string
	X1dEvAllConversions                 string
	X1dEvFirstConversion                string
	X1dPassback                         string
	X1dView                             string
	X1dViewAllConversions               string
	X1dViewFirstConversion              string
	X28dClick                           string
	X28dClickAllConversions             string
	X28dClickFirstConversion            string
	X28dPassback                        string
	X28dView                            string
	X28dViewAllConversions              string
	X28dViewFirstConversion             string
	X7dClick                            string
	X7dClickAllConversions              string
	X7dClickFirstConversion             string
	X7dPassback                         string
	X7dView                             string
	X7dViewAllConversions               string
	X7dViewFirstConversion              string
	ActionBrand                         string
	ActionCanvasComponentID             string
	ActionCanvasComponentName           string
	ActionCarouselCardID                string
	ActionCarouselCardName              string
	ActionCategory                      string
	ActionConvertedProductID            string
	ActionDestination                   string
	ActionDevice                        string
	ActionEventChannel                  string
	ActionLinkClickDestination          string
	ActionLocationCode                  string
	ActionReaction                      string
	ActionTargetID                      string
	ActionType                          string
	ActionVideoAssetID                  string
	ActionVideoSound                    string
	ActionVideoType                     string
	Custom                              string
	Dda                                 string
	Incrementality                      string
	IncrementalityAllConversions        string
	IncrementalityFirstConversion       string
	Inline                              string
	InteractiveComponentStickerID       string
	InteractiveComponentStickerResponse string
	PromotedProductSetResult            string
	SkanClick                           string
	SkanClickSecondPostback             string
	SkanClickThirdPostback              string
	SkanView                            string
	SkanViewSecondPostback              string
	SkanViewThirdPostback               string
	Value                               string
}{
	X1dClick:                            "1d_click",
	X1dClickAllConversions:              "1d_click_all_conversions",
	X1dClickFirstConversion:             "1d_click_first_conversion",
	X1dEv:                               "1d_ev",
	X1dEvAllConversions:                 "1d_ev_all_conversions",
	X1dEvFirstConversion:                "1d_ev_first_conversion",
	X1dPassback:                         "1d_passback",
	X1dView:                             "1d_view",
	X1dViewAllConversions:               "1d_view_all_conversions",
	X1dViewFirstConversion:              "1d_view_first_conversion",
	X28dClick:                           "28d_click",
	X28dClickAllConversions:             "28d_click_all_conversions",
	X28dClickFirstConversion:            "28d_click_first_conversion",
	X28dPassback:                        "28d_passback",
	X28dView:                            "28d_view",
	X28dViewAllConversions:              "28d_view_all_conversions",
	X28dViewFirstConversion:             "28d_view_first_conversion",
	X7dClick:                            "7d_click",
	X7dClickAllConversions:              "7d_click_all_conversions",
	X7dClickFirstConversion:             "7d_click_first_conversion",
	X7dPassback:                         "7d_passback",
	X7dView:                             "7d_view",
	X7dViewAllConversions:               "7d_view_all_conversions",
	X7dViewFirstConversion:              "7d_view_first_conversion",
	ActionBrand:                         "action_brand",
	ActionCanvasComponentID:             "action_canvas_component_id",
	ActionCanvasComponentName:           "action_canvas_component_name",
	ActionCarouselCardID:                "action_carousel_card_id",
	ActionCarouselCardName:              "action_carousel_card_name",
	ActionCategory:                      "action_category",
	ActionConvertedProductID:            "action_converted_product_id",
	ActionDestination:                   "action_destination",
	ActionDevice:                        "action_device",
	ActionEventChannel:                  "action_event_channel",
	ActionLinkClickDestination:          "action_link_click_destination",
	ActionLocationCode:                  "action_location_code",
	ActionReaction:                      "action_reaction",
	ActionTargetID:                      "action_target_id",
	ActionType:                          "action_type",
	ActionVideoAssetID:                  "action_video_asset_id",
	ActionVideoSound:                    "action_video_sound",
	ActionVideoType:                     "action_video_type",
	Custom:                              "custom",
	Dda:                                 "dda",
	Incrementality:                      "incrementality",
	IncrementalityAllConversions:        "incrementality_all_conversions",
	IncrementalityFirstConversion:       "incrementality_first_conversion",
	Inline:                              "inline",
	InteractiveComponentStickerID:       "interactive_component_sticker_id",
	InteractiveComponentStickerResponse: "interactive_component_sticker_response",
	PromotedProductSetResult:            "promoted_product_set_result",
	SkanClick:                           "skan_click",
	SkanClickSecondPostback:             "skan_click_second_postback",
	SkanClickThirdPostback:              "skan_click_third_postback",
	SkanView:                            "skan_view",
	SkanViewSecondPostback:              "skan_view_second_postback",
	SkanViewThirdPostback:               "skan_view_third_postback",
	Value:                               "value",
}
