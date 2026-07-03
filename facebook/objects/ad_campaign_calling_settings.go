package objects

type AdCampaignCallingSettings struct {
	CallForwarding *string `json:"call_forwarding,omitempty"`
	CallTranscript *string `json:"call_transcript,omitempty"`
}
