package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudyObjectiveOffsiteDatasets struct {
	EventNames *[]string `json:"event_names,omitempty"`
	ID         *core.ID  `json:"id,omitempty"`
}
