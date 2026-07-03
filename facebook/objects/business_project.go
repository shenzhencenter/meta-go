package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type BusinessProject struct {
	Business    *Business               `json:"business,omitempty"`
	CreatedTime *time.Time              `json:"created_time,omitempty"`
	Creator     *map[string]interface{} `json:"creator,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
}
