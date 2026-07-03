package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AnalyticsSegment struct {
	CustomAudienceIneligiblityReasons *[]string                 `json:"custom_audience_ineligiblity_reasons,omitempty"`
	Description                       *string                   `json:"description,omitempty"`
	EstimatedCustomAudienceSize       *uint64                   `json:"estimated_custom_audience_size,omitempty"`
	EventInfoRules                    *[]map[string]interface{} `json:"event_info_rules,omitempty"`
	EventRules                        *[]map[string]interface{} `json:"event_rules,omitempty"`
	FilterSet                         *string                   `json:"filter_set,omitempty"`
	HasDemographicRules               *bool                     `json:"has_demographic_rules,omitempty"`
	ID                                *core.ID                  `json:"id,omitempty"`
	IsAllUser                         *bool                     `json:"is_all_user,omitempty"`
	IsEligibleForPushCampaign         *bool                     `json:"is_eligible_for_push_campaign,omitempty"`
	IsInternal                        *bool                     `json:"is_internal,omitempty"`
	Name                              *string                   `json:"name,omitempty"`
	PercentileRules                   *[]map[string]interface{} `json:"percentile_rules,omitempty"`
	TimeLastSeen                      *uint64                   `json:"time_last_seen,omitempty"`
	TimeLastUpdated                   *uint64                   `json:"time_last_updated,omitempty"`
	UserPropertyRules                 *[]map[string]interface{} `json:"user_property_rules,omitempty"`
	WebParamRules                     *[]map[string]interface{} `json:"web_param_rules,omitempty"`
}

var AnalyticsSegmentFields = struct {
	CustomAudienceIneligiblityReasons string
	Description                       string
	EstimatedCustomAudienceSize       string
	EventInfoRules                    string
	EventRules                        string
	FilterSet                         string
	HasDemographicRules               string
	ID                                string
	IsAllUser                         string
	IsEligibleForPushCampaign         string
	IsInternal                        string
	Name                              string
	PercentileRules                   string
	TimeLastSeen                      string
	TimeLastUpdated                   string
	UserPropertyRules                 string
	WebParamRules                     string
}{
	CustomAudienceIneligiblityReasons: "custom_audience_ineligiblity_reasons",
	Description:                       "description",
	EstimatedCustomAudienceSize:       "estimated_custom_audience_size",
	EventInfoRules:                    "event_info_rules",
	EventRules:                        "event_rules",
	FilterSet:                         "filter_set",
	HasDemographicRules:               "has_demographic_rules",
	ID:                                "id",
	IsAllUser:                         "is_all_user",
	IsEligibleForPushCampaign:         "is_eligible_for_push_campaign",
	IsInternal:                        "is_internal",
	Name:                              "name",
	PercentileRules:                   "percentile_rules",
	TimeLastSeen:                      "time_last_seen",
	TimeLastUpdated:                   "time_last_updated",
	UserPropertyRules:                 "user_property_rules",
	WebParamRules:                     "web_param_rules",
}
