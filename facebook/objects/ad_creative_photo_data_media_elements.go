package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativePhotoDataMediaElements struct {
	ElementID   *core.ID `json:"element_id,omitempty"`
	ElementType *string  `json:"element_type,omitempty"`
	X           *float64 `json:"x,omitempty"`
	Y           *float64 `json:"y,omitempty"`
}
