package enums

type ShadowiguserinsightsMetricEnumParam string

const (
	ShadowiguserinsightsMetricEnumParamAccountsEngaged             ShadowiguserinsightsMetricEnumParam = "accounts_engaged"
	ShadowiguserinsightsMetricEnumParamComments                    ShadowiguserinsightsMetricEnumParam = "comments"
	ShadowiguserinsightsMetricEnumParamContentViews                ShadowiguserinsightsMetricEnumParam = "content_views"
	ShadowiguserinsightsMetricEnumParamEngagedAudienceDemographics ShadowiguserinsightsMetricEnumParam = "engaged_audience_demographics"
	ShadowiguserinsightsMetricEnumParamFollowerCount               ShadowiguserinsightsMetricEnumParam = "follower_count"
	ShadowiguserinsightsMetricEnumParamFollowerDemographics        ShadowiguserinsightsMetricEnumParam = "follower_demographics"
	ShadowiguserinsightsMetricEnumParamFollowsAndUnfollows         ShadowiguserinsightsMetricEnumParam = "follows_and_unfollows"
	ShadowiguserinsightsMetricEnumParamLikes                       ShadowiguserinsightsMetricEnumParam = "likes"
	ShadowiguserinsightsMetricEnumParamOnlineFollowers             ShadowiguserinsightsMetricEnumParam = "online_followers"
	ShadowiguserinsightsMetricEnumParamProfileLinksTaps            ShadowiguserinsightsMetricEnumParam = "profile_links_taps"
	ShadowiguserinsightsMetricEnumParamProfileViews                ShadowiguserinsightsMetricEnumParam = "profile_views"
	ShadowiguserinsightsMetricEnumParamQuotes                      ShadowiguserinsightsMetricEnumParam = "quotes"
	ShadowiguserinsightsMetricEnumParamReach                       ShadowiguserinsightsMetricEnumParam = "reach"
	ShadowiguserinsightsMetricEnumParamReachedAudienceDemographics ShadowiguserinsightsMetricEnumParam = "reached_audience_demographics"
	ShadowiguserinsightsMetricEnumParamReplies                     ShadowiguserinsightsMetricEnumParam = "replies"
	ShadowiguserinsightsMetricEnumParamReposts                     ShadowiguserinsightsMetricEnumParam = "reposts"
	ShadowiguserinsightsMetricEnumParamSaves                       ShadowiguserinsightsMetricEnumParam = "saves"
	ShadowiguserinsightsMetricEnumParamShares                      ShadowiguserinsightsMetricEnumParam = "shares"
	ShadowiguserinsightsMetricEnumParamThreadsClicks               ShadowiguserinsightsMetricEnumParam = "threads_clicks"
	ShadowiguserinsightsMetricEnumParamThreadsFollowerDemographics ShadowiguserinsightsMetricEnumParam = "threads_follower_demographics"
	ShadowiguserinsightsMetricEnumParamThreadsFollowers            ShadowiguserinsightsMetricEnumParam = "threads_followers"
	ShadowiguserinsightsMetricEnumParamThreadsLikes                ShadowiguserinsightsMetricEnumParam = "threads_likes"
	ShadowiguserinsightsMetricEnumParamThreadsReplies              ShadowiguserinsightsMetricEnumParam = "threads_replies"
	ShadowiguserinsightsMetricEnumParamThreadsReposts              ShadowiguserinsightsMetricEnumParam = "threads_reposts"
	ShadowiguserinsightsMetricEnumParamThreadsViews                ShadowiguserinsightsMetricEnumParam = "threads_views"
	ShadowiguserinsightsMetricEnumParamTotalInteractions           ShadowiguserinsightsMetricEnumParam = "total_interactions"
	ShadowiguserinsightsMetricEnumParamViews                       ShadowiguserinsightsMetricEnumParam = "views"
	ShadowiguserinsightsMetricEnumParamWebsiteClicks               ShadowiguserinsightsMetricEnumParam = "website_clicks"
)

func (value ShadowiguserinsightsMetricEnumParam) String() string {
	return string(value)
}
