package objects

type AdAssetFeedAdditionalDataPageNudgeMessage struct {
	Enabled      *bool                     `json:"enabled,omitempty"`
	QuickReplies *[]map[string]interface{} `json:"quick_replies,omitempty"`
	Text         *string                   `json:"text,omitempty"`
}

var AdAssetFeedAdditionalDataPageNudgeMessageFields = struct {
	Enabled      string
	QuickReplies string
	Text         string
}{
	Enabled:      "enabled",
	QuickReplies: "quick_replies",
	Text:         "text",
}
