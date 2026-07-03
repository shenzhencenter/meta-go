package objects

type VideoCopyrightSegment struct {
	DurationInSec  *float64 `json:"duration_in_sec,omitempty"`
	MediaType      *string  `json:"media_type,omitempty"`
	StartTimeInSec *float64 `json:"start_time_in_sec,omitempty"`
}
