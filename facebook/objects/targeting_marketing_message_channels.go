package objects

type TargetingMarketingMessageChannels struct {
	Whatsapp *IDName `json:"whatsapp,omitempty"`
}

var TargetingMarketingMessageChannelsFields = struct {
	Whatsapp string
}{
	Whatsapp: "whatsapp",
}
