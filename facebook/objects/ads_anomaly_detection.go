package objects

type AdsAnomalyDetection struct {
	AnomalyData *[]map[string]interface{} `json:"anomaly_data,omitempty"`
	Day         *uint64                   `json:"day,omitempty"`
}
