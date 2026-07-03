package objects

type AdAssetFeedSpecLinkURL struct {
	Adlabels           *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	AndroidURL         *string                      `json:"android_url,omitempty"`
	CarouselSeeMoreURL *string                      `json:"carousel_see_more_url,omitempty"`
	DeeplinkURL        *string                      `json:"deeplink_url,omitempty"`
	DisplayURL         *string                      `json:"display_url,omitempty"`
	IosURL             *string                      `json:"ios_url,omitempty"`
	ObjectStoreUrls    *[]string                    `json:"object_store_urls,omitempty"`
	URLTags            *string                      `json:"url_tags,omitempty"`
	WebsiteURL         *string                      `json:"website_url,omitempty"`
}
