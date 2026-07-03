package enums

type MediainsightsMetricEnumParam string

const (
	MediainsightsMetricEnumParamComments                  MediainsightsMetricEnumParam = "comments"
	MediainsightsMetricEnumParamContentViews              MediainsightsMetricEnumParam = "content_views"
	MediainsightsMetricEnumParamCrosspostedViews          MediainsightsMetricEnumParam = "crossposted_views"
	MediainsightsMetricEnumParamFacebookViews             MediainsightsMetricEnumParam = "facebook_views"
	MediainsightsMetricEnumParamFollows                   MediainsightsMetricEnumParam = "follows"
	MediainsightsMetricEnumParamIgReelsAvgWatchTime       MediainsightsMetricEnumParam = "ig_reels_avg_watch_time"
	MediainsightsMetricEnumParamIgReelsVideoViewTotalTime MediainsightsMetricEnumParam = "ig_reels_video_view_total_time"
	MediainsightsMetricEnumParamImpressions               MediainsightsMetricEnumParam = "impressions"
	MediainsightsMetricEnumParamLikes                     MediainsightsMetricEnumParam = "likes"
	MediainsightsMetricEnumParamLinkClicks                MediainsightsMetricEnumParam = "link_clicks"
	MediainsightsMetricEnumParamNavigation                MediainsightsMetricEnumParam = "navigation"
	MediainsightsMetricEnumParamProfileActivity           MediainsightsMetricEnumParam = "profile_activity"
	MediainsightsMetricEnumParamProfileVisits             MediainsightsMetricEnumParam = "profile_visits"
	MediainsightsMetricEnumParamQuotes                    MediainsightsMetricEnumParam = "quotes"
	MediainsightsMetricEnumParamReach                     MediainsightsMetricEnumParam = "reach"
	MediainsightsMetricEnumParamReelsSkipRate             MediainsightsMetricEnumParam = "reels_skip_rate"
	MediainsightsMetricEnumParamReplies                   MediainsightsMetricEnumParam = "replies"
	MediainsightsMetricEnumParamReposts                   MediainsightsMetricEnumParam = "reposts"
	MediainsightsMetricEnumParamSaved                     MediainsightsMetricEnumParam = "saved"
	MediainsightsMetricEnumParamShares                    MediainsightsMetricEnumParam = "shares"
	MediainsightsMetricEnumParamThreadReplies             MediainsightsMetricEnumParam = "thread_replies"
	MediainsightsMetricEnumParamThreadShares              MediainsightsMetricEnumParam = "thread_shares"
	MediainsightsMetricEnumParamThreadsMediaClicks        MediainsightsMetricEnumParam = "threads_media_clicks"
	MediainsightsMetricEnumParamThreadsReposts            MediainsightsMetricEnumParam = "threads_reposts"
	MediainsightsMetricEnumParamThreadsViews              MediainsightsMetricEnumParam = "threads_views"
	MediainsightsMetricEnumParamTotalComments             MediainsightsMetricEnumParam = "total_comments"
	MediainsightsMetricEnumParamTotalInteractions         MediainsightsMetricEnumParam = "total_interactions"
	MediainsightsMetricEnumParamTotalLikes                MediainsightsMetricEnumParam = "total_likes"
	MediainsightsMetricEnumParamTotalViews                MediainsightsMetricEnumParam = "total_views"
	MediainsightsMetricEnumParamViews                     MediainsightsMetricEnumParam = "views"
)

func (value MediainsightsMetricEnumParam) String() string {
	return string(value)
}
