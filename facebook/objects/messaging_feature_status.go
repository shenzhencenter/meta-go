package objects

type MessagingFeatureStatus struct {
	HopV2        *bool `json:"hop_v2,omitempty"`
	IgMultiApp   *bool `json:"ig_multi_app,omitempty"`
	MsgrMultiApp *bool `json:"msgr_multi_app,omitempty"`
}

var MessagingFeatureStatusFields = struct {
	HopV2        string
	IgMultiApp   string
	MsgrMultiApp string
}{
	HopV2:        "hop_v2",
	IgMultiApp:   "ig_multi_app",
	MsgrMultiApp: "msgr_multi_app",
}
