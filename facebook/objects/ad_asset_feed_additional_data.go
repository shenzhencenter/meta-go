package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetFeedAdditionalData struct {
	AutomatedProductTags           *bool    `json:"automated_product_tags,omitempty"`
	BrandPageID                    *core.ID `json:"brand_page_id,omitempty"`
	IsClickToMessage               *bool    `json:"is_click_to_message,omitempty"`
	MultiShareEndCard              *bool    `json:"multi_share_end_card,omitempty"`
	PageWelcomeMessage             *string  `json:"page_welcome_message,omitempty"`
	PartnerAppWelcomeMessageFlowID *core.ID `json:"partner_app_welcome_message_flow_id,omitempty"`
}
