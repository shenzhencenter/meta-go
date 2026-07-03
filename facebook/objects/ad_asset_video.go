package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetVideo struct {
	CaptionIds      *[]core.ID `json:"caption_ids,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	SourceImageURL  *string    `json:"source_image_url,omitempty"`
	Tag             *string    `json:"tag,omitempty"`
	ThumbnailHash   *string    `json:"thumbnail_hash,omitempty"`
	ThumbnailSource *string    `json:"thumbnail_source,omitempty"`
	ThumbnailURL    *string    `json:"thumbnail_url,omitempty"`
	URL             *string    `json:"url,omitempty"`
	URLTags         *string    `json:"url_tags,omitempty"`
	VideoID         *core.ID   `json:"video_id,omitempty"`
	VideoName       *string    `json:"video_name,omitempty"`
}

var AdAssetVideoFields = struct {
	CaptionIds      string
	ID              string
	SourceImageURL  string
	Tag             string
	ThumbnailHash   string
	ThumbnailSource string
	ThumbnailURL    string
	URL             string
	URLTags         string
	VideoID         string
	VideoName       string
}{
	CaptionIds:      "caption_ids",
	ID:              "id",
	SourceImageURL:  "source_image_url",
	Tag:             "tag",
	ThumbnailHash:   "thumbnail_hash",
	ThumbnailSource: "thumbnail_source",
	ThumbnailURL:    "thumbnail_url",
	URL:             "url",
	URLTags:         "url_tags",
	VideoID:         "video_id",
	VideoName:       "video_name",
}
