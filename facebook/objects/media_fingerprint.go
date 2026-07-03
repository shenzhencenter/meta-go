package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
