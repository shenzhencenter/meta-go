package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCampaignGroupIncrementalConversionOptimizationConfig struct {
	ActionType       *string                   `json:"action_type,omitempty"`
	AdStudyEndTime   *core.Time                `json:"ad_study_end_time,omitempty"`
	AdStudyID        *core.ID                  `json:"ad_study_id,omitempty"`
	AdStudyName      *string                   `json:"ad_study_name,omitempty"`
	AdStudyStartTime *core.Time                `json:"ad_study_start_time,omitempty"`
	CellID           *core.ID                  `json:"cell_id,omitempty"`
	CellName         *string                   `json:"cell_name,omitempty"`
	HoldoutSize      *float64                  `json:"holdout_size,omitempty"`
	IcoType          *string                   `json:"ico_type,omitempty"`
	Objectives       *[]map[string]interface{} `json:"objectives,omitempty"`
}

var AdCampaignGroupIncrementalConversionOptimizationConfigFields = struct {
	ActionType       string
	AdStudyEndTime   string
	AdStudyID        string
	AdStudyName      string
	AdStudyStartTime string
	CellID           string
	CellName         string
	HoldoutSize      string
	IcoType          string
	Objectives       string
}{
	ActionType:       "action_type",
	AdStudyEndTime:   "ad_study_end_time",
	AdStudyID:        "ad_study_id",
	AdStudyName:      "ad_study_name",
	AdStudyStartTime: "ad_study_start_time",
	CellID:           "cell_id",
	CellName:         "cell_name",
	HoldoutSize:      "holdout_size",
	IcoType:          "ico_type",
	Objectives:       "objectives",
}
