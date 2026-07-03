package objects

type TrackingAndConversionWithDefaults struct {
	CustomConversion  *[]map[string]interface{} `json:"custom_conversion,omitempty"`
	CustomTracking    *[]map[string]interface{} `json:"custom_tracking,omitempty"`
	DefaultConversion *[]map[string]interface{} `json:"default_conversion,omitempty"`
	DefaultTracking   *[]map[string]interface{} `json:"default_tracking,omitempty"`
}
