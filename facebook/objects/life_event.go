package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type LifeEvent struct {
	Description *string    `json:"description,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	From        *Page      `json:"from,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	IsHidden    *bool      `json:"is_hidden,omitempty"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	Title       *string    `json:"title,omitempty"`
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}
