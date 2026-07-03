package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsSignalDiagnosticIssue struct {
	DataSourceID                 *AdsPixel  `json:"data_source_id,omitempty"`
	DataSourceType               *string    `json:"data_source_type,omitempty"`
	DiagnosticType               *string    `json:"diagnostic_type,omitempty"`
	EventName                    *string    `json:"event_name,omitempty"`
	TrafficAnomalyDropPercentage *float64   `json:"traffic_anomaly_drop_percentage,omitempty"`
	TrafficAnomalyDropTimestamp  *core.Time `json:"traffic_anomaly_drop_timestamp,omitempty"`
}

var AdsSignalDiagnosticIssueFields = struct {
	DataSourceID                 string
	DataSourceType               string
	DiagnosticType               string
	EventName                    string
	TrafficAnomalyDropPercentage string
	TrafficAnomalyDropTimestamp  string
}{
	DataSourceID:                 "data_source_id",
	DataSourceType:               "data_source_type",
	DiagnosticType:               "diagnostic_type",
	EventName:                    "event_name",
	TrafficAnomalyDropPercentage: "traffic_anomaly_drop_percentage",
	TrafficAnomalyDropTimestamp:  "traffic_anomaly_drop_timestamp",
}
