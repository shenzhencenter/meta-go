package enums

type StoriesinsightsMetricEnumParam string

const (
	StoriesinsightsMetricEnumParamPagesFbStoryReplies                    StoriesinsightsMetricEnumParam = "PAGES_FB_STORY_REPLIES"
	StoriesinsightsMetricEnumParamPagesFbStoryShares                     StoriesinsightsMetricEnumParam = "PAGES_FB_STORY_SHARES"
	StoriesinsightsMetricEnumParamPagesFbStoryStickerInteractions        StoriesinsightsMetricEnumParam = "PAGES_FB_STORY_STICKER_INTERACTIONS"
	StoriesinsightsMetricEnumParamPagesFbStoryThreadLightweightReactions StoriesinsightsMetricEnumParam = "PAGES_FB_STORY_THREAD_LIGHTWEIGHT_REACTIONS"
	StoriesinsightsMetricEnumParamPageStoryImpressionsByStoryID          StoriesinsightsMetricEnumParam = "PAGE_STORY_IMPRESSIONS_BY_STORY_ID"
	StoriesinsightsMetricEnumParamPageStoryImpressionsByStoryIDUnique    StoriesinsightsMetricEnumParam = "PAGE_STORY_IMPRESSIONS_BY_STORY_ID_UNIQUE"
	StoriesinsightsMetricEnumParamStoryInteraction                       StoriesinsightsMetricEnumParam = "STORY_INTERACTION"
	StoriesinsightsMetricEnumParamStoryMediaView                         StoriesinsightsMetricEnumParam = "STORY_MEDIA_VIEW"
	StoriesinsightsMetricEnumParamStoryTotalMediaViewUnique              StoriesinsightsMetricEnumParam = "STORY_TOTAL_MEDIA_VIEW_UNIQUE"
)

func (value StoriesinsightsMetricEnumParam) String() string {
	return string(value)
}
