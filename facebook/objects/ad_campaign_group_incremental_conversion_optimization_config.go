package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdCampaignGroupIncrementalConversionOptimizationConfig struct {
	ActionType       *string                   `json:"action_type,omitempty"`
	AdStudyEndTime   *time.Time                `json:"ad_study_end_time,omitempty"`
	AdStudyID        *core.ID                  `json:"ad_study_id,omitempty"`
	AdStudyName      *string                   `json:"ad_study_name,omitempty"`
	AdStudyStartTime *time.Time                `json:"ad_study_start_time,omitempty"`
	CellID           *core.ID                  `json:"cell_id,omitempty"`
	CellName         *string                   `json:"cell_name,omitempty"`
	HoldoutSize      *float64                  `json:"holdout_size,omitempty"`
	IcoType          *string                   `json:"ico_type,omitempty"`
	Objectives       *[]map[string]interface{} `json:"objectives,omitempty"`
}
