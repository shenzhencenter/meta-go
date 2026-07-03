package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeCollectionThumbnailInfo struct {
	ElementChildIndex *int           `json:"element_child_index,omitempty"`
	ElementCrops      *AdsImageCrops `json:"element_crops,omitempty"`
	ElementID         *core.ID       `json:"element_id,omitempty"`
}
