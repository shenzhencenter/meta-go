package objects

type MessagingFeatureStatus struct {
	HopV2        *bool `json:"hop_v2,omitempty"`
	IgMultiApp   *bool `json:"ig_multi_app,omitempty"`
	MsgrMultiApp *bool `json:"msgr_multi_app,omitempty"`
}
