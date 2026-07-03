package objects

import (
	"time"
)

type AdsSignalDiagnosticIssue struct {
	DataSourceID                 *AdsPixel  `json:"data_source_id,omitempty"`
	DataSourceType               *string    `json:"data_source_type,omitempty"`
	DiagnosticType               *string    `json:"diagnostic_type,omitempty"`
	EventName                    *string    `json:"event_name,omitempty"`
	TrafficAnomalyDropPercentage *float64   `json:"traffic_anomaly_drop_percentage,omitempty"`
	TrafficAnomalyDropTimestamp  *time.Time `json:"traffic_anomaly_drop_timestamp,omitempty"`
}
