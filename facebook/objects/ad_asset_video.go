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
