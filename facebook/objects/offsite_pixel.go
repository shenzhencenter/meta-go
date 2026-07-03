package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type OffsitePixel struct {
	Creator        *string    `json:"creator,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	JsPixel        *string    `json:"js_pixel,omitempty"`
	LastFiringTime *time.Time `json:"last_firing_time,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Tag            *string    `json:"tag,omitempty"`
}
