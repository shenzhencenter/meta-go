package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativePhotoDataMediaElements struct {
	ElementID   *core.ID `json:"element_id,omitempty"`
	ElementType *string  `json:"element_type,omitempty"`
	X           *float64 `json:"x,omitempty"`
	Y           *float64 `json:"y,omitempty"`
}

var AdCreativePhotoDataMediaElementsFields = struct {
	ElementID   string
	ElementType string
	X           string
	Y           string
}{
	ElementID:   "element_id",
	ElementType: "element_type",
	X:           "x",
	Y:           "y",
}
