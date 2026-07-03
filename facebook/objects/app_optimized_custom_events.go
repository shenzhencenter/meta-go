package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppOptimizedCustomEvents struct {
	AppID      *core.ID  `json:"app_id,omitempty"`
	AppName    *string   `json:"app_name,omitempty"`
	EventNames *[]string `json:"event_names,omitempty"`
}

var AppOptimizedCustomEventsFields = struct {
	AppID      string
	AppName    string
	EventNames string
}{
	AppID:      "app_id",
	AppName:    "app_name",
	EventNames: "event_names",
}
