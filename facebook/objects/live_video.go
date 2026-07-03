package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type LiveVideo struct {
	AdBreakConfig              *LiveVideoAdBreakConfig              `json:"ad_break_config,omitempty"`
	AdBreakFailureReason       *string                              `json:"ad_break_failure_reason,omitempty"`
	BroadcastStartTime         *time.Time                           `json:"broadcast_start_time,omitempty"`
	Copyright                  *VideoCopyright                      `json:"copyright,omitempty"`
	CreationTime               *time.Time                           `json:"creation_time,omitempty"`
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
	PlannedStartTime           *time.Time                           `json:"planned_start_time,omitempty"`
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
