package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PagePostExperiment struct {
	AutoResolveSettings          *map[string]interface{}                           `json:"auto_resolve_settings,omitempty"`
	ControlVideoID               *core.ID                                          `json:"control_video_id,omitempty"`
	CreationTime                 *time.Time                                        `json:"creation_time,omitempty"`
	Creator                      *User                                             `json:"creator,omitempty"`
	DeclaredWinningTime          *time.Time                                        `json:"declared_winning_time,omitempty"`
	DeclaredWinningVideoID       *core.ID                                          `json:"declared_winning_video_id,omitempty"`
	Description                  *string                                           `json:"description,omitempty"`
	ExperimentVideoIds           *[]core.ID                                        `json:"experiment_video_ids,omitempty"`
	ID                           *core.ID                                          `json:"id,omitempty"`
	InsightSnapshots             *[]map[string][]map[string]map[string]interface{} `json:"insight_snapshots,omitempty"`
	Name                         *string                                           `json:"name,omitempty"`
	OptimizationGoal             *string                                           `json:"optimization_goal,omitempty"`
	PublishStatus                *string                                           `json:"publish_status,omitempty"`
	PublishTime                  *time.Time                                        `json:"publish_time,omitempty"`
	ScheduledExperimentTimestamp *time.Time                                        `json:"scheduled_experiment_timestamp,omitempty"`
	UpdatedTime                  *time.Time                                        `json:"updated_time,omitempty"`
}
