package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MediaFingerprint struct {
	DurationInSec          *float64                `json:"duration_in_sec,omitempty"`
	FingerprintContentType *string                 `json:"fingerprint_content_type,omitempty"`
	FingerprintType        *string                 `json:"fingerprint_type,omitempty"`
	ID                     *core.ID                `json:"id,omitempty"`
	Metadata               *map[string]interface{} `json:"metadata,omitempty"`
	Title                  *string                 `json:"title,omitempty"`
	UniversalContentID     *core.ID                `json:"universal_content_id,omitempty"`
}

var MediaFingerprintFields = struct {
	DurationInSec          string
	FingerprintContentType string
	FingerprintType        string
	ID                     string
	Metadata               string
	Title                  string
	UniversalContentID     string
}{
	DurationInSec:          "duration_in_sec",
	FingerprintContentType: "fingerprint_content_type",
	FingerprintType:        "fingerprint_type",
	ID:                     "id",
	Metadata:               "metadata",
	Title:                  "title",
	UniversalContentID:     "universal_content_id",
}
