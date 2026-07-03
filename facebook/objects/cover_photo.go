package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CoverPhoto struct {
	CoverID *core.ID `json:"cover_id,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
	OffsetX *float64 `json:"offset_x,omitempty"`
	OffsetY *float64 `json:"offset_y,omitempty"`
	Source  *string  `json:"source,omitempty"`
}

var CoverPhotoFields = struct {
	CoverID string
	ID      string
	OffsetX string
	OffsetY string
	Source  string
}{
	CoverID: "cover_id",
	ID:      "id",
	OffsetX: "offset_x",
	OffsetY: "offset_y",
	Source:  "source",
}
