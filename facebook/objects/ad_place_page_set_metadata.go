package objects

type AdPlacePageSetMetadata struct {
	Audience    *map[string]interface{} `json:"audience,omitempty"`
	Custom      *map[string]interface{} `json:"custom,omitempty"`
	ExtraData   *string                 `json:"extra_data,omitempty"`
	FixedRadius *map[string]interface{} `json:"fixed_radius,omitempty"`
}

var AdPlacePageSetMetadataFields = struct {
	Audience    string
	Custom      string
	ExtraData   string
	FixedRadius string
}{
	Audience:    "audience",
	Custom:      "custom",
	ExtraData:   "extra_data",
	FixedRadius: "fixed_radius",
}
