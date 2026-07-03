package enums

type UserinsightsMetricEnumParam string

const (
	UserinsightsMetricEnumParamAccountsEngaged             UserinsightsMetricEnumParam = "accounts_engaged"
	UserinsightsMetricEnumParamComments                    UserinsightsMetricEnumParam = "comments"
	UserinsightsMetricEnumParamContentViews                UserinsightsMetricEnumParam = "content_views"
	UserinsightsMetricEnumParamEngagedAudienceDemographics UserinsightsMetricEnumParam = "engaged_audience_demographics"
	UserinsightsMetricEnumParamFollowerCount               UserinsightsMetricEnumParam = "follower_count"
	UserinsightsMetricEnumParamFollowerDemographics        UserinsightsMetricEnumParam = "follower_demographics"
	UserinsightsMetricEnumParamFollowsAndUnfollows         UserinsightsMetricEnumParam = "follows_and_unfollows"
	UserinsightsMetricEnumParamLikes                       UserinsightsMetricEnumParam = "likes"
	UserinsightsMetricEnumParamOnlineFollowers             UserinsightsMetricEnumParam = "online_followers"
	UserinsightsMetricEnumParamProfileLinksTaps            UserinsightsMetricEnumParam = "profile_links_taps"
	UserinsightsMetricEnumParamProfileViews                UserinsightsMetricEnumParam = "profile_views"
	UserinsightsMetricEnumParamQuotes                      UserinsightsMetricEnumParam = "quotes"
	UserinsightsMetricEnumParamReach                       UserinsightsMetricEnumParam = "reach"
	UserinsightsMetricEnumParamReachedAudienceDemographics UserinsightsMetricEnumParam = "reached_audience_demographics"
	UserinsightsMetricEnumParamReplies                     UserinsightsMetricEnumParam = "replies"
	UserinsightsMetricEnumParamReposts                     UserinsightsMetricEnumParam = "reposts"
	UserinsightsMetricEnumParamSaves                       UserinsightsMetricEnumParam = "saves"
	UserinsightsMetricEnumParamShares                      UserinsightsMetricEnumParam = "shares"
	UserinsightsMetricEnumParamThreadsClicks               UserinsightsMetricEnumParam = "threads_clicks"
	UserinsightsMetricEnumParamThreadsFollowerDemographics UserinsightsMetricEnumParam = "threads_follower_demographics"
	UserinsightsMetricEnumParamThreadsFollowers            UserinsightsMetricEnumParam = "threads_followers"
	UserinsightsMetricEnumParamThreadsLikes                UserinsightsMetricEnumParam = "threads_likes"
	UserinsightsMetricEnumParamThreadsReplies              UserinsightsMetricEnumParam = "threads_replies"
	UserinsightsMetricEnumParamThreadsReposts              UserinsightsMetricEnumParam = "threads_reposts"
	UserinsightsMetricEnumParamThreadsViews                UserinsightsMetricEnumParam = "threads_views"
	UserinsightsMetricEnumParamTotalInteractions           UserinsightsMetricEnumParam = "total_interactions"
	UserinsightsMetricEnumParamViews                       UserinsightsMetricEnumParam = "views"
	UserinsightsMetricEnumParamWebsiteClicks               UserinsightsMetricEnumParam = "website_clicks"
)

func (value UserinsightsMetricEnumParam) String() string {
	return string(value)
}
