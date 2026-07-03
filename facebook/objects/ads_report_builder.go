package objects

type AdsReportBuilder struct {
	Headers             *map[string]interface{}   `json:"headers,omitempty"`
	Rows                *[]map[string]interface{} `json:"rows,omitempty"`
	SkanReadinessStatus *[]map[string]string      `json:"skan_readiness_status,omitempty"`
}

var AdsReportBuilderFields = struct {
	Headers             string
	Rows                string
	SkanReadinessStatus string
}{
	Headers:             "headers",
	Rows:                "rows",
	SkanReadinessStatus: "skan_readiness_status",
}
