package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PagePostExperiment struct {
	AutoResolveSettings          *map[string]interface{}                           `json:"auto_resolve_settings,omitempty"`
	ControlVideoID               *core.ID                                          `json:"control_video_id,omitempty"`
	CreationTime                 *core.Time                                        `json:"creation_time,omitempty"`
	Creator                      *User                                             `json:"creator,omitempty"`
	DeclaredWinningTime          *core.Time                                        `json:"declared_winning_time,omitempty"`
	DeclaredWinningVideoID       *core.ID                                          `json:"declared_winning_video_id,omitempty"`
	Description                  *string                                           `json:"description,omitempty"`
	ExperimentVideoIds           *[]core.ID                                        `json:"experiment_video_ids,omitempty"`
	ID                           *core.ID                                          `json:"id,omitempty"`
	InsightSnapshots             *[]map[string][]map[string]map[string]interface{} `json:"insight_snapshots,omitempty"`
	Name                         *string                                           `json:"name,omitempty"`
	OptimizationGoal             *string                                           `json:"optimization_goal,omitempty"`
	PublishStatus                *string                                           `json:"publish_status,omitempty"`
	PublishTime                  *core.Time                                        `json:"publish_time,omitempty"`
	ScheduledExperimentTimestamp *core.Time                                        `json:"scheduled_experiment_timestamp,omitempty"`
	UpdatedTime                  *core.Time                                        `json:"updated_time,omitempty"`
}

var PagePostExperimentFields = struct {
	AutoResolveSettings          string
	ControlVideoID               string
	CreationTime                 string
	Creator                      string
	DeclaredWinningTime          string
	DeclaredWinningVideoID       string
	Description                  string
	ExperimentVideoIds           string
	ID                           string
	InsightSnapshots             string
	Name                         string
	OptimizationGoal             string
	PublishStatus                string
	PublishTime                  string
	ScheduledExperimentTimestamp string
	UpdatedTime                  string
}{
	AutoResolveSettings:          "auto_resolve_settings",
	ControlVideoID:               "control_video_id",
	CreationTime:                 "creation_time",
	Creator:                      "creator",
	DeclaredWinningTime:          "declared_winning_time",
	DeclaredWinningVideoID:       "declared_winning_video_id",
	Description:                  "description",
	ExperimentVideoIds:           "experiment_video_ids",
	ID:                           "id",
	InsightSnapshots:             "insight_snapshots",
	Name:                         "name",
	OptimizationGoal:             "optimization_goal",
	PublishStatus:                "publish_status",
	PublishTime:                  "publish_time",
	ScheduledExperimentTimestamp: "scheduled_experiment_timestamp",
	UpdatedTime:                  "updated_time",
}
