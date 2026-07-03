package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudyObjectiveOffsiteDatasets struct {
	EventNames *[]string `json:"event_names,omitempty"`
	ID         *core.ID  `json:"id,omitempty"`
}

var AdStudyObjectiveOffsiteDatasetsFields = struct {
	EventNames string
	ID         string
}{
	EventNames: "event_names",
	ID:         "id",
}
