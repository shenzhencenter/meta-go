package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CoverPhoto struct {
	CoverID *core.ID `json:"cover_id,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
	OffsetX *float64 `json:"offset_x,omitempty"`
	OffsetY *float64 `json:"offset_y,omitempty"`
	Source  *string  `json:"source,omitempty"`
}
