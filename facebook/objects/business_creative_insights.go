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

var BusinessCreativeInsightsFields = struct {
	Actions                     string
	Age                         string
	Country                     string
	DateEnd                     string
	DateStart                   string
	DevicePlatform              string
	Gender                      string
	Impressions                 string
	InlineLinkClicks            string
	Objective                   string
	OptimizationGoal            string
	PlatformPosition            string
	PublisherPlatform           string
	QualityRanking              string
	VideoPlayActions            string
	VideoThruplayWatchedActions string
}{
	Actions:                     "actions",
	Age:                         "age",
	Country:                     "country",
	DateEnd:                     "date_end",
	DateStart:                   "date_start",
	DevicePlatform:              "device_platform",
	Gender:                      "gender",
	Impressions:                 "impressions",
	InlineLinkClicks:            "inline_link_clicks",
	Objective:                   "objective",
	OptimizationGoal:            "optimization_goal",
	PlatformPosition:            "platform_position",
	PublisherPlatform:           "publisher_platform",
	QualityRanking:              "quality_ranking",
	VideoPlayActions:            "video_play_actions",
	VideoThruplayWatchedActions: "video_thruplay_watched_actions",
}
