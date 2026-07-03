package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type ThirdPartyPartnerPanelScheduled struct {
	AdentitiesIds   *[]core.ID                                      `json:"adentities_ids,omitempty"`
	Cadence         *string                                         `json:"cadence,omitempty"`
	Country         *string                                         `json:"country,omitempty"`
	CreatedTime     *time.Time                                      `json:"created_time,omitempty"`
	Description     *string                                         `json:"description,omitempty"`
	EndTime         *time.Time                                      `json:"end_time,omitempty"`
	ID              *core.ID                                        `json:"id,omitempty"`
	ModifiedTime    *time.Time                                      `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                        `json:"owner_instance_id,omitempty"`
	OwnerPanelID    *core.ID                                        `json:"owner_panel_id,omitempty"`
	OwnerPanelName  *string                                         `json:"owner_panel_name,omitempty"`
	StartTime       *time.Time                                      `json:"start_time,omitempty"`
	Status          *enums.ThirdpartypartnerpanelscheduledStatus    `json:"status,omitempty"`
	StudyType       *enums.ThirdpartypartnerpanelscheduledStudyType `json:"study_type,omitempty"`
}
