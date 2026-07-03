package objects

type BusinessCreativeInsights struct {
	Actions                     *[]AdsActionStats `json:"actions,omitempty"`
	Age                         *string           `json:"age,omitempty"`
	Country                     *string           `json:"country,omitempty"`
	DateEnd                     *string           `json:"date_end,omitempty"`
	DateStart                   *string           `json:"date_start,omitempty"`
	DevicePlatform              *string           `json:"device_platform,omitempty"`
	Gender                      *string           `json:"gender,omitempty"`
	Impressions                 *int              `json:"impressions,omitempty"`
	InlineLinkClicks            *int              `json:"inline_link_clicks,omitempty"`
	Objective                   *string           `json:"objective,omitempty"`
	OptimizationGoal            *string           `json:"optimization_goal,omitempty"`
	PlatformPosition            *string           `json:"platform_position,omitempty"`
	PublisherPlatform           *string           `json:"publisher_platform,omitempty"`
	QualityRanking              *string           `json:"quality_ranking,omitempty"`
	VideoPlayActions            *[]AdsActionStats `json:"video_play_actions,omitempty"`
	VideoThruplayWatchedActions *[]AdsActionStats `json:"video_thruplay_watched_actions,omitempty"`
}
