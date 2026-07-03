package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ThirdPartyPartnerPanelScheduled struct {
	AdentitiesIds   *[]core.ID                                      `json:"adentities_ids,omitempty"`
	Cadence         *string                                         `json:"cadence,omitempty"`
	Country         *string                                         `json:"country,omitempty"`
	CreatedTime     *core.Time                                      `json:"created_time,omitempty"`
	Description     *string                                         `json:"description,omitempty"`
	EndTime         *core.Time                                      `json:"end_time,omitempty"`
	ID              *core.ID                                        `json:"id,omitempty"`
	ModifiedTime    *core.Time                                      `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                        `json:"owner_instance_id,omitempty"`
	OwnerPanelID    *core.ID                                        `json:"owner_panel_id,omitempty"`
	OwnerPanelName  *string                                         `json:"owner_panel_name,omitempty"`
	StartTime       *core.Time                                      `json:"start_time,omitempty"`
	Status          *enums.ThirdpartypartnerpanelscheduledStatus    `json:"status,omitempty"`
	StudyType       *enums.ThirdpartypartnerpanelscheduledStudyType `json:"study_type,omitempty"`
}

var ThirdPartyPartnerPanelScheduledFields = struct {
	AdentitiesIds   string
	Cadence         string
	Country         string
	CreatedTime     string
	Description     string
	EndTime         string
	ID              string
	ModifiedTime    string
	OwnerInstanceID string
	OwnerPanelID    string
	OwnerPanelName  string
	StartTime       string
	Status          string
	StudyType       string
}{
	AdentitiesIds:   "adentities_ids",
	Cadence:         "cadence",
	Country:         "country",
	CreatedTime:     "created_time",
	Description:     "description",
	EndTime:         "end_time",
	ID:              "id",
	ModifiedTime:    "modified_time",
	OwnerInstanceID: "owner_instance_id",
	OwnerPanelID:    "owner_panel_id",
	OwnerPanelName:  "owner_panel_name",
	StartTime:       "start_time",
	Status:          "status",
	StudyType:       "study_type",
}
