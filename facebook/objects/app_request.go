package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppRequest struct {
	ActionType  *string                 `json:"action_type,omitempty"`
	Application *Application            `json:"application,omitempty"`
	CreatedTime *core.Time              `json:"created_time,omitempty"`
	Data        *string                 `json:"data,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Message     *string                 `json:"message,omitempty"`
	Object      *map[string]interface{} `json:"object,omitempty"`
	To          *map[string]interface{} `json:"to,omitempty"`
}

var AppRequestFields = struct {
	ActionType  string
	Application string
	CreatedTime string
	Data        string
	From        string
	ID          string
	Message     string
	Object      string
	To          string
}{
	ActionType:  "action_type",
	Application: "application",
	CreatedTime: "created_time",
	Data:        "data",
	From:        "from",
	ID:          "id",
	Message:     "message",
	Object:      "object",
	To:          "to",
}
