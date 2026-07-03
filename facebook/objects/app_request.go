package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AppRequest struct {
	ActionType  *string                 `json:"action_type,omitempty"`
	Application *Application            `json:"application,omitempty"`
	CreatedTime *time.Time              `json:"created_time,omitempty"`
	Data        *string                 `json:"data,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Message     *string                 `json:"message,omitempty"`
	Object      *map[string]interface{} `json:"object,omitempty"`
	To          *map[string]interface{} `json:"to,omitempty"`
}
