package objects

type AdsPixelEventLastFiredTime struct {
	Event         *string `json:"event,omitempty"`
	LastFiredTime *int    `json:"last_fired_time,omitempty"`
}

var AdsPixelEventLastFiredTimeFields = struct {
	Event         string
	LastFiredTime string
}{
	Event:         "event",
	LastFiredTime: "last_fired_time",
}
