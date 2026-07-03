package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsStartYourDayWidget struct {
	ID       *core.ID `json:"id,omitempty"`
	WidgetID *core.ID `json:"widget_id,omitempty"`
}

var AdsStartYourDayWidgetFields = struct {
	ID       string
	WidgetID string
}{
	ID:       "id",
	WidgetID: "widget_id",
}
