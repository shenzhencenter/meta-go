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

var AdAssetFeedSpecLinkURLFields = struct {
	Adlabels           string
	AndroidURL         string
	CarouselSeeMoreURL string
	DeeplinkURL        string
	DisplayURL         string
	IosURL             string
	ObjectStoreUrls    string
	URLTags            string
	WebsiteURL         string
}{
	Adlabels:           "adlabels",
	AndroidURL:         "android_url",
	CarouselSeeMoreURL: "carousel_see_more_url",
	DeeplinkURL:        "deeplink_url",
	DisplayURL:         "display_url",
	IosURL:             "ios_url",
	ObjectStoreUrls:    "object_store_urls",
	URLTags:            "url_tags",
	WebsiteURL:         "website_url",
}
