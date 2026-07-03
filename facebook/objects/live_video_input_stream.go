package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LiveVideoInputStream struct {
	DashIngestURL   *string                 `json:"dash_ingest_url,omitempty"`
	DashPreviewURL  *string                 `json:"dash_preview_url,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
	IsMaster        *bool                   `json:"is_master,omitempty"`
	SecureStreamURL *string                 `json:"secure_stream_url,omitempty"`
	StreamHealth    *map[string]interface{} `json:"stream_health,omitempty"`
	StreamID        *core.ID                `json:"stream_id,omitempty"`
	StreamURL       *string                 `json:"stream_url,omitempty"`
}

var LiveVideoInputStreamFields = struct {
	DashIngestURL   string
	DashPreviewURL  string
	ID              string
	IsMaster        string
	SecureStreamURL string
	StreamHealth    string
	StreamID        string
	StreamURL       string
}{
	DashIngestURL:   "dash_ingest_url",
	DashPreviewURL:  "dash_preview_url",
	ID:              "id",
	IsMaster:        "is_master",
	SecureStreamURL: "secure_stream_url",
	StreamHealth:    "stream_health",
	StreamID:        "stream_id",
	StreamURL:       "stream_url",
}
