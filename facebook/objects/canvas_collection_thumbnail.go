package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CanvasCollectionThumbnail struct {
	ElementChildIndex *int     `json:"element_child_index,omitempty"`
	ElementID         *core.ID `json:"element_id,omitempty"`
	Photo             *Photo   `json:"photo,omitempty"`
}

var CanvasCollectionThumbnailFields = struct {
	ElementChildIndex string
	ElementID         string
	Photo             string
}{
	ElementChildIndex: "element_child_index",
	ElementID:         "element_id",
	Photo:             "photo",
}
