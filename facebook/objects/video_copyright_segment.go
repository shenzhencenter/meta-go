package objects

type VideoCopyrightSegment struct {
	DurationInSec  *float64 `json:"duration_in_sec,omitempty"`
	MediaType      *string  `json:"media_type,omitempty"`
	StartTimeInSec *float64 `json:"start_time_in_sec,omitempty"`
}

var VideoCopyrightSegmentFields = struct {
	DurationInSec  string
	MediaType      string
	StartTimeInSec string
}{
	DurationInSec:  "duration_in_sec",
	MediaType:      "media_type",
	StartTimeInSec: "start_time_in_sec",
}
