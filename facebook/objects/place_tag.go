package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type PlaceTag struct {
	CreatedTime *time.Time `json:"created_time,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Place       *Page      `json:"place,omitempty"`
}
