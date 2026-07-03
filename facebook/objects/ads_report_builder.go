package objects

type AdsReportBuilder struct {
	Headers             *map[string]interface{}   `json:"headers,omitempty"`
	Rows                *[]map[string]interface{} `json:"rows,omitempty"`
	SkanReadinessStatus *[]map[string]string      `json:"skan_readiness_status,omitempty"`
}
