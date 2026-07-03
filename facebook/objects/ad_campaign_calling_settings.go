package objects

type AdCampaignCallingSettings struct {
	CallForwarding *string `json:"call_forwarding,omitempty"`
	CallTranscript *string `json:"call_transcript,omitempty"`
}

var AdCampaignCallingSettingsFields = struct {
	CallForwarding string
	CallTranscript string
}{
	CallForwarding: "call_forwarding",
	CallTranscript: "call_transcript",
}
