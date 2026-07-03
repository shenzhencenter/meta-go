package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CopyrightReferenceContainer struct {
	ContentType           *string                 `json:"content_type,omitempty"`
	CopyrightCreationTime *time.Time              `json:"copyright_creation_time,omitempty"`
	DownloadHdURL         *string                 `json:"download_hd_url,omitempty"`
	DurationInSec         *float64                `json:"duration_in_sec,omitempty"`
	ID                    *core.ID                `json:"id,omitempty"`
	Iswc                  *string                 `json:"iswc,omitempty"`
	Metadata              *map[string]interface{} `json:"metadata,omitempty"`
	PlayableVideoURI      *string                 `json:"playable_video_uri,omitempty"`
	PublishedTime         *time.Time              `json:"published_time,omitempty"`
	ThumbnailURL          *string                 `json:"thumbnail_url,omitempty"`
	Title                 *string                 `json:"title,omitempty"`
	UniversalContentID    *core.ID                `json:"universal_content_id,omitempty"`
	WriterNames           *[]string               `json:"writer_names,omitempty"`
}
