package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeCollectionThumbnailInfo struct {
	ElementChildIndex *int           `json:"element_child_index,omitempty"`
	ElementCrops      *AdsImageCrops `json:"element_crops,omitempty"`
	ElementID         *core.ID       `json:"element_id,omitempty"`
}

var AdCreativeCollectionThumbnailInfoFields = struct {
	ElementChildIndex string
	ElementCrops      string
	ElementID         string
}{
	ElementChildIndex: "element_child_index",
	ElementCrops:      "element_crops",
	ElementID:         "element_id",
}
