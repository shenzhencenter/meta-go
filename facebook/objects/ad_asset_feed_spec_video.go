package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetFeedSpecVideo struct {
	Adlabels      *[]AdAssetFeedSpecAssetLabel `json:"adlabels,omitempty"`
	CaptionIds    *[]core.ID                   `json:"caption_ids,omitempty"`
	ThumbnailHash *string                      `json:"thumbnail_hash,omitempty"`
	ThumbnailURL  *string                      `json:"thumbnail_url,omitempty"`
	URLTags       *string                      `json:"url_tags,omitempty"`
	VideoID       *core.ID                     `json:"video_id,omitempty"`
}

var AdAssetFeedSpecVideoFields = struct {
	Adlabels      string
	CaptionIds    string
	ThumbnailHash string
	ThumbnailURL  string
	URLTags       string
	VideoID       string
}{
	Adlabels:      "adlabels",
	CaptionIds:    "caption_ids",
	ThumbnailHash: "thumbnail_hash",
	ThumbnailURL:  "thumbnail_url",
	URLTags:       "url_tags",
	VideoID:       "video_id",
}
