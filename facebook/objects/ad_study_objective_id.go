package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudyObjectiveID struct {
	EventNames *[]string `json:"event_names,omitempty"`
	ID         *core.ID  `json:"id,omitempty"`
	Type       *string   `json:"type,omitempty"`
}
