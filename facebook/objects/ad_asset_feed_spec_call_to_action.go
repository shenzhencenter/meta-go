package objects

type AdAssetFeedSpecCallToAction struct {
	Adlabels *[]AdAssetFeedSpecAssetLabel         `json:"adlabels,omitempty"`
	Type     *string                              `json:"type,omitempty"`
	Value    *AdCreativeLinkDataCallToActionValue `json:"value,omitempty"`
}

var AdAssetFeedSpecCallToActionFields = struct {
	Adlabels string
	Type     string
	Value    string
}{
	Adlabels: "adlabels",
	Type:     "type",
	Value:    "value",
}
