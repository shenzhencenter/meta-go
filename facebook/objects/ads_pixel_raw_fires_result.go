package objects

import (
	"time"
)

type AdsPixelRawFiresResult struct {
	DataJSON             *string    `json:"data_json,omitempty"`
	DeviceType           *string    `json:"device_type,omitempty"`
	Event                *string    `json:"event,omitempty"`
	EventDetectionMethod *string    `json:"event_detection_method,omitempty"`
	EventSrc             *string    `json:"event_src,omitempty"`
	PlacedURL            *string    `json:"placed_url,omitempty"`
	Timestamp            *time.Time `json:"timestamp,omitempty"`
	UserPiiKeys          *string    `json:"user_pii_keys,omitempty"`
}
