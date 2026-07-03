package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetImage struct {
	Hash       *string        `json:"hash,omitempty"`
	ID         *core.ID       `json:"id,omitempty"`
	ImageCrops *AdsImageCrops `json:"image_crops,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Tag        *string        `json:"tag,omitempty"`
	URL        *string        `json:"url,omitempty"`
	URLTags    *string        `json:"url_tags,omitempty"`
}
