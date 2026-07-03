package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ThirdPartyPartnerPanelRequest struct {
	AdentitiesIds   *[]core.ID                                    `json:"adentities_ids,omitempty"`
	Country         *string                                       `json:"country,omitempty"`
	CreatedTime     *core.Time                                    `json:"created_time,omitempty"`
	Description     *string                                       `json:"description,omitempty"`
	ID              *core.ID                                      `json:"id,omitempty"`
	ModifiedTime    *core.Time                                    `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                      `json:"owner_instance_id,omitempty"`
	OwnerPanelID    *core.ID                                      `json:"owner_panel_id,omitempty"`
	OwnerPanelName  *string                                       `json:"owner_panel_name,omitempty"`
	Status          *enums.ThirdpartypartnerpanelrequestStatus    `json:"status,omitempty"`
	StudyEndTime    *core.Time                                    `json:"study_end_time,omitempty"`
	StudyStartTime  *core.Time                                    `json:"study_start_time,omitempty"`
	StudyType       *enums.ThirdpartypartnerpanelrequestStudyType `json:"study_type,omitempty"`
}

var ThirdPartyPartnerPanelRequestFields = struct {
	AdentitiesIds   string
	Country         string
	CreatedTime     string
	Description     string
	ID              string
	ModifiedTime    string
	OwnerInstanceID string
	OwnerPanelID    string
	OwnerPanelName  string
	Status          string
	StudyEndTime    string
	StudyStartTime  string
	StudyType       string
}{
	AdentitiesIds:   "adentities_ids",
	Country:         "country",
	CreatedTime:     "created_time",
	Description:     "description",
	ID:              "id",
	ModifiedTime:    "modified_time",
	OwnerInstanceID: "owner_instance_id",
	OwnerPanelID:    "owner_panel_id",
	OwnerPanelName:  "owner_panel_name",
	Status:          "status",
	StudyEndTime:    "study_end_time",
	StudyStartTime:  "study_start_time",
	StudyType:       "study_type",
}
