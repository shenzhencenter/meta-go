package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
