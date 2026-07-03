package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAssetFeedSpecVideo struct {
	Adlabels      *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	CaptionIds    *[]core.ID                   `json:"caption_ids,omitempty"`
	ThumbnailHash *string                      `json:"thumbnail_hash,omitempty"`
	ThumbnailURL  *string                      `json:"thumbnail_url,omitempty"`
	URLTags       *string                      `json:"url_tags,omitempty"`
	VideoID       *core.ID                     `json:"video_id,omitempty"`
}
