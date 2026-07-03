package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelRawFiresResult struct {
	DataJSON             *string    `json:"data_json,omitempty"`
	DeviceType           *string    `json:"device_type,omitempty"`
	Event                *string    `json:"event,omitempty"`
	EventDetectionMethod *string    `json:"event_detection_method,omitempty"`
	EventSrc             *string    `json:"event_src,omitempty"`
	PlacedURL            *string    `json:"placed_url,omitempty"`
	Timestamp            *core.Time `json:"timestamp,omitempty"`
	UserPiiKeys          *string    `json:"user_pii_keys,omitempty"`
}

var AdsPixelRawFiresResultFields = struct {
	DataJSON             string
	DeviceType           string
	Event                string
	EventDetectionMethod string
	EventSrc             string
	PlacedURL            string
	Timestamp            string
	UserPiiKeys          string
}{
	DataJSON:             "data_json",
	DeviceType:           "device_type",
	Event:                "event",
	EventDetectionMethod: "event_detection_method",
	EventSrc:             "event_src",
	PlacedURL:            "placed_url",
	Timestamp:            "timestamp",
	UserPiiKeys:          "user_pii_keys",
}
