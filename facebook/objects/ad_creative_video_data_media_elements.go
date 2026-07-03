package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeVideoDataMediaElements struct {
	ElementID   *core.ID `json:"element_id,omitempty"`
	ElementType *string  `json:"element_type,omitempty"`
}
