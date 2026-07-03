package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsHistogramStats struct {
	X1dClick                            *[]int   `json:"1d_click,omitempty"`
	X1dClickAllConversions              *[]int   `json:"1d_click_all_conversions,omitempty"`
	X1dClickFirstConversion             *[]int   `json:"1d_click_first_conversion,omitempty"`
	X1dEv                               *[]int   `json:"1d_ev,omitempty"`
	X1dEvAllConversions                 *[]int   `json:"1d_ev_all_conversions,omitempty"`
	X1dEvFirstConversion                *[]int   `json:"1d_ev_first_conversion,omitempty"`
	X1dPassback                         *[]int   `json:"1d_passback,omitempty"`
	X1dView                             *[]int   `json:"1d_view,omitempty"`
	X1dViewAllConversions               *[]int   `json:"1d_view_all_conversions,omitempty"`
	X1dViewFirstConversion              *[]int   `json:"1d_view_first_conversion,omitempty"`
	X28dClick                           *[]int   `json:"28d_click,omitempty"`
	X28dClickAllConversions             *[]int   `json:"28d_click_all_conversions,omitempty"`
	X28dClickFirstConversion            *[]int   `json:"28d_click_first_conversion,omitempty"`
	X28dPassback                        *[]int   `json:"28d_passback,omitempty"`
	X28dView                            *[]int   `json:"28d_view,omitempty"`
	X28dViewAllConversions              *[]int   `json:"28d_view_all_conversions,omitempty"`
	X28dViewFirstConversion             *[]int   `json:"28d_view_first_conversion,omitempty"`
	X7dClick                            *[]int   `json:"7d_click,omitempty"`
	X7dClickAllConversions              *[]int   `json:"7d_click_all_conversions,omitempty"`
	X7dClickFirstConversion             *[]int   `json:"7d_click_first_conversion,omitempty"`
	X7dPassback                         *[]int   `json:"7d_passback,omitempty"`
	X7dView                             *[]int   `json:"7d_view,omitempty"`
	X7dViewAllConversions               *[]int   `json:"7d_view_all_conversions,omitempty"`
	X7dViewFirstConversion              *[]int   `json:"7d_view_first_conversion,omitempty"`
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
	Custom                              *[]int   `json:"custom,omitempty"`
	Dda                                 *[]int   `json:"dda,omitempty"`
	Incrementality                      *[]int   `json:"incrementality,omitempty"`
	IncrementalityAllConversions        *[]int   `json:"incrementality_all_conversions,omitempty"`
	IncrementalityFirstConversion       *[]int   `json:"incrementality_first_conversion,omitempty"`
	Inline                              *[]int   `json:"inline,omitempty"`
	InteractiveComponentStickerID       *core.ID `json:"interactive_component_sticker_id,omitempty"`
	InteractiveComponentStickerResponse *string  `json:"interactive_component_sticker_response,omitempty"`
	PromotedProductSetResult            *string  `json:"promoted_product_set_result,omitempty"`
	SkanClick                           *[]int   `json:"skan_click,omitempty"`
	SkanClickSecondPostback             *[]int   `json:"skan_click_second_postback,omitempty"`
	SkanClickThirdPostback              *[]int   `json:"skan_click_third_postback,omitempty"`
	SkanView                            *[]int   `json:"skan_view,omitempty"`
	SkanViewSecondPostback              *[]int   `json:"skan_view_second_postback,omitempty"`
	SkanViewThirdPostback               *[]int   `json:"skan_view_third_postback,omitempty"`
	Value                               *[]int   `json:"value,omitempty"`
}
