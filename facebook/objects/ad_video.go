package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdVideo struct {
	AdBreaks                  *[]int                    `json:"ad_breaks,omitempty"`
	AdminCreator              *User                     `json:"admin_creator,omitempty"`
	AudioIsrc                 *AudioIsrc                `json:"audio_isrc,omitempty"`
	BackdatedTime             *core.Time                `json:"backdated_time,omitempty"`
	BackdatedTimeGranularity  *string                   `json:"backdated_time_granularity,omitempty"`
	BoostEligibilityInfo      *map[string]interface{}   `json:"boost_eligibility_info,omitempty"`
	ContentCategory           *string                   `json:"content_category,omitempty"`
	ContentTags               *[]string                 `json:"content_tags,omitempty"`
	Copyright                 *VideoCopyright           `json:"copyright,omitempty"`
	CopyrightCheckInformation *map[string]interface{}   `json:"copyright_check_information,omitempty"`
	CopyrightMonitoringStatus *string                   `json:"copyright_monitoring_status,omitempty"`
	CreatedTime               *core.Time                `json:"created_time,omitempty"`
	CustomLabels              *[]string                 `json:"custom_labels,omitempty"`
	Description               *string                   `json:"description,omitempty"`
	EmbedHTML                 *map[string]interface{}   `json:"embed_html,omitempty"`
	Embeddable                *bool                     `json:"embeddable,omitempty"`
	Event                     *Event                    `json:"event,omitempty"`
	Expiration                *map[string]interface{}   `json:"expiration,omitempty"`
	Format                    *[]map[string]interface{} `json:"format,omitempty"`
	From                      *map[string]interface{}   `json:"from,omitempty"`
	Icon                      *string                   `json:"icon,omitempty"`
	ID                        *core.ID                  `json:"id,omitempty"`
	IsCrosspostVideo          *bool                     `json:"is_crosspost_video,omitempty"`
	IsCrosspostingEligible    *bool                     `json:"is_crossposting_eligible,omitempty"`
	IsEpisode                 *bool                     `json:"is_episode,omitempty"`
	IsInstagramEligible       *bool                     `json:"is_instagram_eligible,omitempty"`
	IsReferenceOnly           *bool                     `json:"is_reference_only,omitempty"`
	Length                    *float64                  `json:"length,omitempty"`
	LiveAudienceCount         *uint64                   `json:"live_audience_count,omitempty"`
	LiveStatus                *string                   `json:"live_status,omitempty"`
	MusicVideoCopyright       *MusicVideoCopyright      `json:"music_video_copyright,omitempty"`
	PermalinkURL              *string                   `json:"permalink_url,omitempty"`
	Picture                   *string                   `json:"picture,omitempty"`
	Place                     *Place                    `json:"place,omitempty"`
	PostID                    *core.ID                  `json:"post_id,omitempty"`
	PostViews                 *uint64                   `json:"post_views,omitempty"`
	PremiereLivingRoomStatus  *string                   `json:"premiere_living_room_status,omitempty"`
	Privacy                   *Privacy                  `json:"privacy,omitempty"`
	Published                 *bool                     `json:"published,omitempty"`
	ReplaceAudioStatus        *string                   `json:"replace_audio_status,omitempty"`
	ScheduledPublishTime      *core.Time                `json:"scheduled_publish_time,omitempty"`
	SelectedAudioSpec         *map[string]interface{}   `json:"selected_audio_spec,omitempty"`
	Source                    *string                   `json:"source,omitempty"`
	Spherical                 *bool                     `json:"spherical,omitempty"`
	Status                    *VideoStatus              `json:"status,omitempty"`
	Title                     *string                   `json:"title,omitempty"`
	UniversalVideoID          *core.ID                  `json:"universal_video_id,omitempty"`
	UpdatedTime               *core.Time                `json:"updated_time,omitempty"`
	Views                     *uint64                   `json:"views,omitempty"`
}

var AdVideoFields = struct {
	AdBreaks                  string
	AdminCreator              string
	AudioIsrc                 string
	BackdatedTime             string
	BackdatedTimeGranularity  string
	BoostEligibilityInfo      string
	ContentCategory           string
	ContentTags               string
	Copyright                 string
	CopyrightCheckInformation string
	CopyrightMonitoringStatus string
	CreatedTime               string
	CustomLabels              string
	Description               string
	EmbedHTML                 string
	Embeddable                string
	Event                     string
	Expiration                string
	Format                    string
	From                      string
	Icon                      string
	ID                        string
	IsCrosspostVideo          string
	IsCrosspostingEligible    string
	IsEpisode                 string
	IsInstagramEligible       string
	IsReferenceOnly           string
	Length                    string
	LiveAudienceCount         string
	LiveStatus                string
	MusicVideoCopyright       string
	PermalinkURL              string
	Picture                   string
	Place                     string
	PostID                    string
	PostViews                 string
	PremiereLivingRoomStatus  string
	Privacy                   string
	Published                 string
	ReplaceAudioStatus        string
	ScheduledPublishTime      string
	SelectedAudioSpec         string
	Source                    string
	Spherical                 string
	Status                    string
	Title                     string
	UniversalVideoID          string
	UpdatedTime               string
	Views                     string
}{
	AdBreaks:                  "ad_breaks",
	AdminCreator:              "admin_creator",
	AudioIsrc:                 "audio_isrc",
	BackdatedTime:             "backdated_time",
	BackdatedTimeGranularity:  "backdated_time_granularity",
	BoostEligibilityInfo:      "boost_eligibility_info",
	ContentCategory:           "content_category",
	ContentTags:               "content_tags",
	Copyright:                 "copyright",
	CopyrightCheckInformation: "copyright_check_information",
	CopyrightMonitoringStatus: "copyright_monitoring_status",
	CreatedTime:               "created_time",
	CustomLabels:              "custom_labels",
	Description:               "description",
	EmbedHTML:                 "embed_html",
	Embeddable:                "embeddable",
	Event:                     "event",
	Expiration:                "expiration",
	Format:                    "format",
	From:                      "from",
	Icon:                      "icon",
	ID:                        "id",
	IsCrosspostVideo:          "is_crosspost_video",
	IsCrosspostingEligible:    "is_crossposting_eligible",
	IsEpisode:                 "is_episode",
	IsInstagramEligible:       "is_instagram_eligible",
	IsReferenceOnly:           "is_reference_only",
	Length:                    "length",
	LiveAudienceCount:         "live_audience_count",
	LiveStatus:                "live_status",
	MusicVideoCopyright:       "music_video_copyright",
	PermalinkURL:              "permalink_url",
	Picture:                   "picture",
	Place:                     "place",
	PostID:                    "post_id",
	PostViews:                 "post_views",
	PremiereLivingRoomStatus:  "premiere_living_room_status",
	Privacy:                   "privacy",
	Published:                 "published",
	ReplaceAudioStatus:        "replace_audio_status",
	ScheduledPublishTime:      "scheduled_publish_time",
	SelectedAudioSpec:         "selected_audio_spec",
	Source:                    "source",
	Spherical:                 "spherical",
	Status:                    "status",
	Title:                     "title",
	UniversalVideoID:          "universal_video_id",
	UpdatedTime:               "updated_time",
	Views:                     "views",
}
