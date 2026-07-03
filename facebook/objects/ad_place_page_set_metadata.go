package objects

type AdPlacePageSetMetadata struct {
	Audience    *map[string]interface{} `json:"audience,omitempty"`
	Custom      *map[string]interface{} `json:"custom,omitempty"`
	ExtraData   *string                 `json:"extra_data,omitempty"`
	FixedRadius *map[string]interface{} `json:"fixed_radius,omitempty"`
}
