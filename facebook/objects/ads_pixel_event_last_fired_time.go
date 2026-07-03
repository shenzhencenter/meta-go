package objects

type AdsPixelEventLastFiredTime struct {
	Event         *string `json:"event,omitempty"`
	LastFiredTime *int    `json:"last_fired_time,omitempty"`
}
