package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CanvasCollectionThumbnail struct {
	ElementChildIndex *int     `json:"element_child_index,omitempty"`
	ElementID         *core.ID `json:"element_id,omitempty"`
	Photo             *Photo   `json:"photo,omitempty"`
}
