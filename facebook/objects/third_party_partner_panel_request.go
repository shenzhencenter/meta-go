package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"time"
)

type ThirdPartyPartnerPanelRequest struct {
	AdentitiesIds   *[]core.ID                                    `json:"adentities_ids,omitempty"`
	Country         *string                                       `json:"country,omitempty"`
	CreatedTime     *time.Time                                    `json:"created_time,omitempty"`
	Description     *string                                       `json:"description,omitempty"`
	ID              *core.ID                                      `json:"id,omitempty"`
	ModifiedTime    *time.Time                                    `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                      `json:"owner_instance_id,omitempty"`
	OwnerPanelID    *core.ID                                      `json:"owner_panel_id,omitempty"`
	OwnerPanelName  *string                                       `json:"owner_panel_name,omitempty"`
	Status          *enums.ThirdpartypartnerpanelrequestStatus    `json:"status,omitempty"`
	StudyEndTime    *time.Time                                    `json:"study_end_time,omitempty"`
	StudyStartTime  *time.Time                                    `json:"study_start_time,omitempty"`
	StudyType       *enums.ThirdpartypartnerpanelrequestStudyType `json:"study_type,omitempty"`
}
