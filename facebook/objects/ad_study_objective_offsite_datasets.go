package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdStudyObjectiveOffsiteDatasets struct {
	EventNames *[]string `json:"event_names,omitempty"`
	ID         *core.ID  `json:"id,omitempty"`
}
