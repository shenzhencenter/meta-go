package objects

type BusinessTrafficAnalysisReport struct {
	AudienceLocation           *[]map[string]interface{} `json:"audience_location,omitempty"`
	EventCategory              *[]map[string]interface{} `json:"event_category,omitempty"`
	TrafficAnalysisImpressions *[]map[string]interface{} `json:"traffic_analysis_impressions,omitempty"`
}

var BusinessTrafficAnalysisReportFields = struct {
	AudienceLocation           string
	EventCategory              string
	TrafficAnalysisImpressions string
}{
	AudienceLocation:           "audience_location",
	EventCategory:              "event_category",
	TrafficAnalysisImpressions: "traffic_analysis_impressions",
}
