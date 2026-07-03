package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type IGMedia struct {
	AltText                   *string                                  `json:"alt_text,omitempty"`
	BoostEligibilityInfo      *IGMediaBoostEligibilityInfo             `json:"boost_eligibility_info,omitempty"`
	Caption                   *string                                  `json:"caption,omitempty"`
	CommentsCount             *int                                     `json:"comments_count,omitempty"`
	CopyrightCheckInformation *IGVideoCopyrightCheckMatchesInformation `json:"copyright_check_information,omitempty"`
	CurrentLiveViewerCount    *int                                     `json:"current_live_viewer_count,omitempty"`
	HasPoll                   *bool                                    `json:"has_poll,omitempty"`
	HasSlider                 *bool                                    `json:"has_slider,omitempty"`
	ID                        *core.ID                                 `json:"id,omitempty"`
	IgID                      *core.ID                                 `json:"ig_id,omitempty"`
	IsCommentEnabled          *bool                                    `json:"is_comment_enabled,omitempty"`
	IsSharedToFeed            *bool                                    `json:"is_shared_to_feed,omitempty"`
	LegacyInstagramMediaID    *core.ID                                 `json:"legacy_instagram_media_id,omitempty"`
	LikeCount                 *int                                     `json:"like_count,omitempty"`
	MediaAudioType            *string                                  `json:"media_audio_type,omitempty"`
	MediaProductType          *string                                  `json:"media_product_type,omitempty"`
	MediaType                 *string                                  `json:"media_type,omitempty"`
	MediaURL                  *string                                  `json:"media_url,omitempty"`
	Owner                     *IGUser                                  `json:"owner,omitempty"`
	Permalink                 *string                                  `json:"permalink,omitempty"`
	RepostsCount              *int                                     `json:"reposts_count,omitempty"`
	SavedCount                *int                                     `json:"saved_count,omitempty"`
	SharesCount               *int                                     `json:"shares_count,omitempty"`
	Shortcode                 *string                                  `json:"shortcode,omitempty"`
	ThumbnailURL              *string                                  `json:"thumbnail_url,omitempty"`
	Timestamp                 *time.Time                               `json:"timestamp,omitempty"`
	TotalCommentsCount        *int                                     `json:"total_comments_count,omitempty"`
	TotalLikeCount            *int                                     `json:"total_like_count,omitempty"`
	TotalViewsCount           *int                                     `json:"total_views_count,omitempty"`
	Username                  *string                                  `json:"username,omitempty"`
	VideoTitle                *string                                  `json:"video_title,omitempty"`
	ViewCount                 *int                                     `json:"view_count,omitempty"`
}
