package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdsPixelRealTimeEventLogResult struct {
	DataJSON                   *string  `json:"data_json,omitempty"`
	DedupData                  *string  `json:"dedup_data,omitempty"`
	DeviceType                 *string  `json:"device_type,omitempty"`
	DomainControlRuleRejection *string  `json:"domain_control_rule_rejection,omitempty"`
	Event                      *string  `json:"event,omitempty"`
	EventDetectionMethod       *string  `json:"event_detection_method,omitempty"`
	InIframe                   *bool    `json:"in_iframe,omitempty"`
	MatchedRuleConditions      *string  `json:"matched_rule_conditions,omitempty"`
	ResolvedLink               *string  `json:"resolved_link,omitempty"`
	SourceRuleCondition        *string  `json:"source_rule_condition,omitempty"`
	Timestamp                  *string  `json:"timestamp,omitempty"`
	TraceID                    *core.ID `json:"trace_id,omitempty"`
	URL                        *string  `json:"url,omitempty"`
}
