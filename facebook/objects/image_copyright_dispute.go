package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type ImageCopyrightDispute struct {
	AppealFormData  *string    `json:"appeal_form_data,omitempty"`
	DisputeFormData *string    `json:"dispute_form_data,omitempty"`
	ExpirationTime  *time.Time `json:"expiration_time,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	MatchID         *core.ID   `json:"match_id,omitempty"`
	Status          *string    `json:"status,omitempty"`
	TimeAppealed    *time.Time `json:"time_appealed,omitempty"`
	TimeCreated     *time.Time `json:"time_created,omitempty"`
	TimeUpdated     *time.Time `json:"time_updated,omitempty"`
}
