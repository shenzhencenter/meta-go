package enums

type BusinessvideosValidationAdPlacementsEnumParam string

const (
	BusinessvideosValidationAdPlacementsEnumParamAudienceNetworkInstreamVideo       BusinessvideosValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_INSTREAM_VIDEO"
	BusinessvideosValidationAdPlacementsEnumParamAudienceNetworkInstreamVideoMobile BusinessvideosValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_INSTREAM_VIDEO_MOBILE"
	BusinessvideosValidationAdPlacementsEnumParamAudienceNetworkRewardedVideo       BusinessvideosValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_REWARDED_VIDEO"
	BusinessvideosValidationAdPlacementsEnumParamDesktopFeedStandard                BusinessvideosValidationAdPlacementsEnumParam = "DESKTOP_FEED_STANDARD"
	BusinessvideosValidationAdPlacementsEnumParamFacebookStoryMobile                BusinessvideosValidationAdPlacementsEnumParam = "FACEBOOK_STORY_MOBILE"
	BusinessvideosValidationAdPlacementsEnumParamFacebookStoryStickerMobile         BusinessvideosValidationAdPlacementsEnumParam = "FACEBOOK_STORY_STICKER_MOBILE"
	BusinessvideosValidationAdPlacementsEnumParamInstagramStandard                  BusinessvideosValidationAdPlacementsEnumParam = "INSTAGRAM_STANDARD"
	BusinessvideosValidationAdPlacementsEnumParamInstagramStory                     BusinessvideosValidationAdPlacementsEnumParam = "INSTAGRAM_STORY"
	BusinessvideosValidationAdPlacementsEnumParamInstantArticleStandard             BusinessvideosValidationAdPlacementsEnumParam = "INSTANT_ARTICLE_STANDARD"
	BusinessvideosValidationAdPlacementsEnumParamInstreamBannerDesktop              BusinessvideosValidationAdPlacementsEnumParam = "INSTREAM_BANNER_DESKTOP"
	BusinessvideosValidationAdPlacementsEnumParamInstreamBannerMobile               BusinessvideosValidationAdPlacementsEnumParam = "INSTREAM_BANNER_MOBILE"
	BusinessvideosValidationAdPlacementsEnumParamInstreamVideoDesktop               BusinessvideosValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_DESKTOP"
	BusinessvideosValidationAdPlacementsEnumParamInstreamVideoImage                 BusinessvideosValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_IMAGE"
	BusinessvideosValidationAdPlacementsEnumParamInstreamVideoMobile                BusinessvideosValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_MOBILE"
	BusinessvideosValidationAdPlacementsEnumParamMessengerMobileInboxMedia          BusinessvideosValidationAdPlacementsEnumParam = "MESSENGER_MOBILE_INBOX_MEDIA"
	BusinessvideosValidationAdPlacementsEnumParamMessengerMobileStoryMedia          BusinessvideosValidationAdPlacementsEnumParam = "MESSENGER_MOBILE_STORY_MEDIA"
	BusinessvideosValidationAdPlacementsEnumParamMobileFeedStandard                 BusinessvideosValidationAdPlacementsEnumParam = "MOBILE_FEED_STANDARD"
	BusinessvideosValidationAdPlacementsEnumParamMobileFullwidth                    BusinessvideosValidationAdPlacementsEnumParam = "MOBILE_FULLWIDTH"
	BusinessvideosValidationAdPlacementsEnumParamMobileInterstitial                 BusinessvideosValidationAdPlacementsEnumParam = "MOBILE_INTERSTITIAL"
	BusinessvideosValidationAdPlacementsEnumParamMobileMediumRectangle              BusinessvideosValidationAdPlacementsEnumParam = "MOBILE_MEDIUM_RECTANGLE"
	BusinessvideosValidationAdPlacementsEnumParamMobileNative                       BusinessvideosValidationAdPlacementsEnumParam = "MOBILE_NATIVE"
	BusinessvideosValidationAdPlacementsEnumParamRightColumnStandard                BusinessvideosValidationAdPlacementsEnumParam = "RIGHT_COLUMN_STANDARD"
	BusinessvideosValidationAdPlacementsEnumParamSuggestedVideoMobile               BusinessvideosValidationAdPlacementsEnumParam = "SUGGESTED_VIDEO_MOBILE"
)

func (value BusinessvideosValidationAdPlacementsEnumParam) String() string {
	return string(value)
}
