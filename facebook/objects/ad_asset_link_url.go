package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetLinkURL struct {
	AndroidDeeplinkURL *string  `json:"android_deeplink_url,omitempty"`
	CarouselSeeMoreURL *string  `json:"carousel_see_more_url,omitempty"`
	DeeplinkURL        *string  `json:"deeplink_url,omitempty"`
	DisplayURL         *string  `json:"display_url,omitempty"`
	ID                 *core.ID `json:"id,omitempty"`
	IpadDeeplinkURL    *string  `json:"ipad_deeplink_url,omitempty"`
	IphoneDeeplinkURL  *string  `json:"iphone_deeplink_url,omitempty"`
	URLTags            *string  `json:"url_tags,omitempty"`
	WebsiteURL         *string  `json:"website_url,omitempty"`
}
