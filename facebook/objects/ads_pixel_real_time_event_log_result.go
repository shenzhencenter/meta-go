package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdsPixelRealTimeEventLogResultFields = struct {
	DataJSON                   string
	DedupData                  string
	DeviceType                 string
	DomainControlRuleRejection string
	Event                      string
	EventDetectionMethod       string
	InIframe                   string
	MatchedRuleConditions      string
	ResolvedLink               string
	SourceRuleCondition        string
	Timestamp                  string
	TraceID                    string
	URL                        string
}{
	DataJSON:                   "data_json",
	DedupData:                  "dedup_data",
	DeviceType:                 "device_type",
	DomainControlRuleRejection: "domain_control_rule_rejection",
	Event:                      "event",
	EventDetectionMethod:       "event_detection_method",
	InIframe:                   "in_iframe",
	MatchedRuleConditions:      "matched_rule_conditions",
	ResolvedLink:               "resolved_link",
	SourceRuleCondition:        "source_rule_condition",
	Timestamp:                  "timestamp",
	TraceID:                    "trace_id",
	URL:                        "url",
}
