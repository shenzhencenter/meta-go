package objects

type OfflineConversionDataSetOptimizationStatus struct {
	Event            *string `json:"event,omitempty"`
	LastChangedTime  *int    `json:"last_changed_time,omitempty"`
	LastDetectedTime *int    `json:"last_detected_time,omitempty"`
	Status           *string `json:"status,omitempty"`
}

var OfflineConversionDataSetOptimizationStatusFields = struct {
	Event            string
	LastChangedTime  string
	LastDetectedTime string
	Status           string
}{
	Event:            "event",
	LastChangedTime:  "last_changed_time",
	LastDetectedTime: "last_detected_time",
	Status:           "status",
}
