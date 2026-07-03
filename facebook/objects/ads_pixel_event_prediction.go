package objects

type AdsPixelEventPrediction struct {
	Dismissed *bool   `json:"dismissed,omitempty"`
	EventType *string `json:"event_type,omitempty"`
	Rule      *string `json:"rule,omitempty"`
}

var AdsPixelEventPredictionFields = struct {
	Dismissed string
	EventType string
	Rule      string
}{
	Dismissed: "dismissed",
	EventType: "event_type",
	Rule:      "rule",
}
