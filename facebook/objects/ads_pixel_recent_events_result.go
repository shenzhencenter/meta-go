package objects

type AdsPixelRecentEventsResult struct {
	Count *int    `json:"count,omitempty"`
	Event *string `json:"event,omitempty"`
}

var AdsPixelRecentEventsResultFields = struct {
	Count string
	Event string
}{
	Count: "count",
	Event: "event",
}
