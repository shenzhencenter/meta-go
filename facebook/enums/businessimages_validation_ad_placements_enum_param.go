package enums

type BusinessimagesValidationAdPlacementsEnumParam string

const (
	BusinessimagesValidationAdPlacementsEnumParamAudienceNetworkInstreamVideo       BusinessimagesValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_INSTREAM_VIDEO"
	BusinessimagesValidationAdPlacementsEnumParamAudienceNetworkInstreamVideoMobile BusinessimagesValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_INSTREAM_VIDEO_MOBILE"
	BusinessimagesValidationAdPlacementsEnumParamAudienceNetworkRewardedVideo       BusinessimagesValidationAdPlacementsEnumParam = "AUDIENCE_NETWORK_REWARDED_VIDEO"
	BusinessimagesValidationAdPlacementsEnumParamDesktopFeedStandard                BusinessimagesValidationAdPlacementsEnumParam = "DESKTOP_FEED_STANDARD"
	BusinessimagesValidationAdPlacementsEnumParamFacebookStoryMobile                BusinessimagesValidationAdPlacementsEnumParam = "FACEBOOK_STORY_MOBILE"
	BusinessimagesValidationAdPlacementsEnumParamFacebookStoryStickerMobile         BusinessimagesValidationAdPlacementsEnumParam = "FACEBOOK_STORY_STICKER_MOBILE"
	BusinessimagesValidationAdPlacementsEnumParamInstagramStandard                  BusinessimagesValidationAdPlacementsEnumParam = "INSTAGRAM_STANDARD"
	BusinessimagesValidationAdPlacementsEnumParamInstagramStory                     BusinessimagesValidationAdPlacementsEnumParam = "INSTAGRAM_STORY"
	BusinessimagesValidationAdPlacementsEnumParamInstantArticleStandard             BusinessimagesValidationAdPlacementsEnumParam = "INSTANT_ARTICLE_STANDARD"
	BusinessimagesValidationAdPlacementsEnumParamInstreamBannerDesktop              BusinessimagesValidationAdPlacementsEnumParam = "INSTREAM_BANNER_DESKTOP"
	BusinessimagesValidationAdPlacementsEnumParamInstreamBannerMobile               BusinessimagesValidationAdPlacementsEnumParam = "INSTREAM_BANNER_MOBILE"
	BusinessimagesValidationAdPlacementsEnumParamInstreamVideoDesktop               BusinessimagesValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_DESKTOP"
	BusinessimagesValidationAdPlacementsEnumParamInstreamVideoImage                 BusinessimagesValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_IMAGE"
	BusinessimagesValidationAdPlacementsEnumParamInstreamVideoMobile                BusinessimagesValidationAdPlacementsEnumParam = "INSTREAM_VIDEO_MOBILE"
	BusinessimagesValidationAdPlacementsEnumParamMessengerMobileInboxMedia          BusinessimagesValidationAdPlacementsEnumParam = "MESSENGER_MOBILE_INBOX_MEDIA"
	BusinessimagesValidationAdPlacementsEnumParamMessengerMobileStoryMedia          BusinessimagesValidationAdPlacementsEnumParam = "MESSENGER_MOBILE_STORY_MEDIA"
	BusinessimagesValidationAdPlacementsEnumParamMobileFeedStandard                 BusinessimagesValidationAdPlacementsEnumParam = "MOBILE_FEED_STANDARD"
	BusinessimagesValidationAdPlacementsEnumParamMobileFullwidth                    BusinessimagesValidationAdPlacementsEnumParam = "MOBILE_FULLWIDTH"
	BusinessimagesValidationAdPlacementsEnumParamMobileInterstitial                 BusinessimagesValidationAdPlacementsEnumParam = "MOBILE_INTERSTITIAL"
	BusinessimagesValidationAdPlacementsEnumParamMobileMediumRectangle              BusinessimagesValidationAdPlacementsEnumParam = "MOBILE_MEDIUM_RECTANGLE"
	BusinessimagesValidationAdPlacementsEnumParamMobileNative                       BusinessimagesValidationAdPlacementsEnumParam = "MOBILE_NATIVE"
	BusinessimagesValidationAdPlacementsEnumParamRightColumnStandard                BusinessimagesValidationAdPlacementsEnumParam = "RIGHT_COLUMN_STANDARD"
	BusinessimagesValidationAdPlacementsEnumParamSuggestedVideoMobile               BusinessimagesValidationAdPlacementsEnumParam = "SUGGESTED_VIDEO_MOBILE"
)

func (value BusinessimagesValidationAdPlacementsEnumParam) String() string {
	return string(value)
}
