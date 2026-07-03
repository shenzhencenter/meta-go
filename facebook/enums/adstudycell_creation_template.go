package enums

type AdstudycellCreationTemplate string

const (
	AdstudycellCreationTemplateAutomaticPlacements           AdstudycellCreationTemplate = "AUTOMATIC_PLACEMENTS"
	AdstudycellCreationTemplateBrandAwareness                AdstudycellCreationTemplate = "BRAND_AWARENESS"
	AdstudycellCreationTemplateFacebook                      AdstudycellCreationTemplate = "FACEBOOK"
	AdstudycellCreationTemplateFacebookAudienceNetwork       AdstudycellCreationTemplate = "FACEBOOK_AUDIENCE_NETWORK"
	AdstudycellCreationTemplateFacebookInstagram             AdstudycellCreationTemplate = "FACEBOOK_INSTAGRAM"
	AdstudycellCreationTemplateFacebookNewsFeed              AdstudycellCreationTemplate = "FACEBOOK_NEWS_FEED"
	AdstudycellCreationTemplateFacebookNewsFeedInStreamVideo AdstudycellCreationTemplate = "FACEBOOK_NEWS_FEED_IN_STREAM_VIDEO"
	AdstudycellCreationTemplateHighFrequency                 AdstudycellCreationTemplate = "HIGH_FREQUENCY"
	AdstudycellCreationTemplateInstagram                     AdstudycellCreationTemplate = "INSTAGRAM"
	AdstudycellCreationTemplateInStreamVideo                 AdstudycellCreationTemplate = "IN_STREAM_VIDEO"
	AdstudycellCreationTemplateLowFrequency                  AdstudycellCreationTemplate = "LOW_FREQUENCY"
	AdstudycellCreationTemplateMediumFrequency               AdstudycellCreationTemplate = "MEDIUM_FREQUENCY"
	AdstudycellCreationTemplateMobileOptimizedVideo          AdstudycellCreationTemplate = "MOBILE_OPTIMIZED_VIDEO"
	AdstudycellCreationTemplatePagePostEngagement            AdstudycellCreationTemplate = "PAGE_POST_ENGAGEMENT"
	AdstudycellCreationTemplateReach                         AdstudycellCreationTemplate = "REACH"
	AdstudycellCreationTemplateTvCommercial                  AdstudycellCreationTemplate = "TV_COMMERCIAL"
	AdstudycellCreationTemplateTvFacebook                    AdstudycellCreationTemplate = "TV_FACEBOOK"
	AdstudycellCreationTemplateVideoViewOptimization         AdstudycellCreationTemplate = "VIDEO_VIEW_OPTIMIZATION"
)

func (value AdstudycellCreationTemplate) String() string {
	return string(value)
}
