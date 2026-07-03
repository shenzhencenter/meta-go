package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MediaCopyrightAttribution struct {
	AttributionIgTargetID         *core.ID   `json:"attribution_ig_target_id,omitempty"`
	AttributionTargetEmailAddress *string    `json:"attribution_target_email_address,omitempty"`
	AttributionTargetID           *core.ID   `json:"attribution_target_id,omitempty"`
	AttributionTargetName         *string    `json:"attribution_target_name,omitempty"`
	AttributionType               *string    `json:"attribution_type,omitempty"`
	AttributionURI                *string    `json:"attribution_uri,omitempty"`
	CopyrightCount                *int       `json:"copyright_count,omitempty"`
	CreationTime                  *core.Time `json:"creation_time,omitempty"`
	Creator                       *Profile   `json:"creator,omitempty"`
	ID                            *core.ID   `json:"id,omitempty"`
	IsEnabled                     *bool      `json:"is_enabled,omitempty"`
	LinkTitle                     *string    `json:"link_title,omitempty"`
	MatchCount                    *int       `json:"match_count,omitempty"`
	Status                        *string    `json:"status,omitempty"`
	Title                         *string    `json:"title,omitempty"`
}

var MediaCopyrightAttributionFields = struct {
	AttributionIgTargetID         string
	AttributionTargetEmailAddress string
	AttributionTargetID           string
	AttributionTargetName         string
	AttributionType               string
	AttributionURI                string
	CopyrightCount                string
	CreationTime                  string
	Creator                       string
	ID                            string
	IsEnabled                     string
	LinkTitle                     string
	MatchCount                    string
	Status                        string
	Title                         string
}{
	AttributionIgTargetID:         "attribution_ig_target_id",
	AttributionTargetEmailAddress: "attribution_target_email_address",
	AttributionTargetID:           "attribution_target_id",
	AttributionTargetName:         "attribution_target_name",
	AttributionType:               "attribution_type",
	AttributionURI:                "attribution_uri",
	CopyrightCount:                "copyright_count",
	CreationTime:                  "creation_time",
	Creator:                       "creator",
	ID:                            "id",
	IsEnabled:                     "is_enabled",
	LinkTitle:                     "link_title",
	MatchCount:                    "match_count",
	Status:                        "status",
	Title:                         "title",
}
