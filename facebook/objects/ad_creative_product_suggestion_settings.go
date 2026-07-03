package objects

type AdCreativeProductSuggestionSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
}

var AdCreativeProductSuggestionSettingsFields = struct {
	Enabled string
}{
	Enabled: "enabled",
}
