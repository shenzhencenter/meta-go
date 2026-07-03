package objects

type ReachFrequencySpec struct {
	Countries                      *[]string               `json:"countries,omitempty"`
	DefaultCreationData            *map[string]interface{} `json:"default_creation_data,omitempty"`
	GlobalIoMaxCampaignDuration    *uint64                 `json:"global_io_max_campaign_duration,omitempty"`
	MaxCampaignDuration            *map[string]interface{} `json:"max_campaign_duration,omitempty"`
	MaxDaysToFinish                *map[string]interface{} `json:"max_days_to_finish,omitempty"`
	MaxPauseWithoutPredictionRerun *map[string]interface{} `json:"max_pause_without_prediction_rerun,omitempty"`
	MinCampaignDuration            *map[string]interface{} `json:"min_campaign_duration,omitempty"`
	MinReachLimits                 *map[string]interface{} `json:"min_reach_limits,omitempty"`
}

var ReachFrequencySpecFields = struct {
	Countries                      string
	DefaultCreationData            string
	GlobalIoMaxCampaignDuration    string
	MaxCampaignDuration            string
	MaxDaysToFinish                string
	MaxPauseWithoutPredictionRerun string
	MinCampaignDuration            string
	MinReachLimits                 string
}{
	Countries:                      "countries",
	DefaultCreationData:            "default_creation_data",
	GlobalIoMaxCampaignDuration:    "global_io_max_campaign_duration",
	MaxCampaignDuration:            "max_campaign_duration",
	MaxDaysToFinish:                "max_days_to_finish",
	MaxPauseWithoutPredictionRerun: "max_pause_without_prediction_rerun",
	MinCampaignDuration:            "min_campaign_duration",
	MinReachLimits:                 "min_reach_limits",
}
