package objects

type OfflineConversionDataSetOptimizationStatus struct {
	Event            *string `json:"event,omitempty"`
	LastChangedTime  *int    `json:"last_changed_time,omitempty"`
	LastDetectedTime *int    `json:"last_detected_time,omitempty"`
	Status           *string `json:"status,omitempty"`
}
