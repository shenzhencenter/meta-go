package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LiveVideo struct {
	AdBreakConfig              *LiveVideoAdBreakConfig              `json:"ad_break_config,omitempty"`
	AdBreakFailureReason       *string                              `json:"ad_break_failure_reason,omitempty"`
	BroadcastStartTime         *core.Time                           `json:"broadcast_start_time,omitempty"`
	Copyright                  *VideoCopyright                      `json:"copyright,omitempty"`
	CreationTime               *core.Time                           `json:"creation_time,omitempty"`
	DashIngestURL              *string                              `json:"dash_ingest_url,omitempty"`
	DashPreviewURL             *string                              `json:"dash_preview_url,omitempty"`
	Description                *string                              `json:"description,omitempty"`
	EmbedHTML                  *map[string]interface{}              `json:"embed_html,omitempty"`
	From                       *map[string]interface{}              `json:"from,omitempty"`
	ID                         *core.ID                             `json:"id,omitempty"`
	IngestStreams              *[]LiveVideoInputStream              `json:"ingest_streams,omitempty"`
	IsManualMode               *bool                                `json:"is_manual_mode,omitempty"`
	IsReferenceOnly            *bool                                `json:"is_reference_only,omitempty"`
	LiveViews                  *uint64                              `json:"live_views,omitempty"`
	PermalinkURL               *string                              `json:"permalink_url,omitempty"`
	PlannedStartTime           *core.Time                           `json:"planned_start_time,omitempty"`
	RecommendedEncoderSettings *LiveVideoRecommendedEncoderSettings `json:"recommended_encoder_settings,omitempty"`
	SecondsLeft                *int                                 `json:"seconds_left,omitempty"`
	SecureStreamURL            *string                              `json:"secure_stream_url,omitempty"`
	Status                     *string                              `json:"status,omitempty"`
	StreamURL                  *string                              `json:"stream_url,omitempty"`
	Targeting                  *LiveVideoTargeting                  `json:"targeting,omitempty"`
	Title                      *string                              `json:"title,omitempty"`
	TotalViews                 *string                              `json:"total_views,omitempty"`
	Video                      *AdVideo                             `json:"video,omitempty"`
}

var LiveVideoFields = struct {
	AdBreakConfig              string
	AdBreakFailureReason       string
	BroadcastStartTime         string
	Copyright                  string
	CreationTime               string
	DashIngestURL              string
	DashPreviewURL             string
	Description                string
	EmbedHTML                  string
	From                       string
	ID                         string
	IngestStreams              string
	IsManualMode               string
	IsReferenceOnly            string
	LiveViews                  string
	PermalinkURL               string
	PlannedStartTime           string
	RecommendedEncoderSettings string
	SecondsLeft                string
	SecureStreamURL            string
	Status                     string
	StreamURL                  string
	Targeting                  string
	Title                      string
	TotalViews                 string
	Video                      string
}{
	AdBreakConfig:              "ad_break_config",
	AdBreakFailureReason:       "ad_break_failure_reason",
	BroadcastStartTime:         "broadcast_start_time",
	Copyright:                  "copyright",
	CreationTime:               "creation_time",
	DashIngestURL:              "dash_ingest_url",
	DashPreviewURL:             "dash_preview_url",
	Description:                "description",
	EmbedHTML:                  "embed_html",
	From:                       "from",
	ID:                         "id",
	IngestStreams:              "ingest_streams",
	IsManualMode:               "is_manual_mode",
	IsReferenceOnly:            "is_reference_only",
	LiveViews:                  "live_views",
	PermalinkURL:               "permalink_url",
	PlannedStartTime:           "planned_start_time",
	RecommendedEncoderSettings: "recommended_encoder_settings",
	SecondsLeft:                "seconds_left",
	SecureStreamURL:            "secure_stream_url",
	Status:                     "status",
	StreamURL:                  "stream_url",
	Targeting:                  "targeting",
	Title:                      "title",
	TotalViews:                 "total_views",
	Video:                      "video",
}
