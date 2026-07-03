package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type Status struct {
	Event       *Event                  `json:"event,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Message     *string                 `json:"message,omitempty"`
	Place       *Place                  `json:"place,omitempty"`
	UpdatedTime *time.Time              `json:"updated_time,omitempty"`
}
