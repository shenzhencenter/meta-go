package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ImageCopyrightDispute struct {
	AppealFormData  *string    `json:"appeal_form_data,omitempty"`
	DisputeFormData *string    `json:"dispute_form_data,omitempty"`
	ExpirationTime  *core.Time `json:"expiration_time,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	MatchID         *core.ID   `json:"match_id,omitempty"`
	Status          *string    `json:"status,omitempty"`
	TimeAppealed    *core.Time `json:"time_appealed,omitempty"`
	TimeCreated     *core.Time `json:"time_created,omitempty"`
	TimeUpdated     *core.Time `json:"time_updated,omitempty"`
}

var ImageCopyrightDisputeFields = struct {
	AppealFormData  string
	DisputeFormData string
	ExpirationTime  string
	ID              string
	MatchID         string
	Status          string
	TimeAppealed    string
	TimeCreated     string
	TimeUpdated     string
}{
	AppealFormData:  "appeal_form_data",
	DisputeFormData: "dispute_form_data",
	ExpirationTime:  "expiration_time",
	ID:              "id",
	MatchID:         "match_id",
	Status:          "status",
	TimeAppealed:    "time_appealed",
	TimeCreated:     "time_created",
	TimeUpdated:     "time_updated",
}
