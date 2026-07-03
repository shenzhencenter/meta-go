package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type MediaCopyrightAttribution struct {
	AttributionIgTargetID         *core.ID   `json:"attribution_ig_target_id,omitempty"`
	AttributionTargetEmailAddress *string    `json:"attribution_target_email_address,omitempty"`
	AttributionTargetID           *core.ID   `json:"attribution_target_id,omitempty"`
	AttributionTargetName         *string    `json:"attribution_target_name,omitempty"`
	AttributionType               *string    `json:"attribution_type,omitempty"`
	AttributionURI                *string    `json:"attribution_uri,omitempty"`
	CopyrightCount                *int       `json:"copyright_count,omitempty"`
	CreationTime                  *time.Time `json:"creation_time,omitempty"`
	Creator                       *Profile   `json:"creator,omitempty"`
	ID                            *core.ID   `json:"id,omitempty"`
	IsEnabled                     *bool      `json:"is_enabled,omitempty"`
	LinkTitle                     *string    `json:"link_title,omitempty"`
	MatchCount                    *int       `json:"match_count,omitempty"`
	Status                        *string    `json:"status,omitempty"`
	Title                         *string    `json:"title,omitempty"`
}
