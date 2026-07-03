package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CopyrightReferenceContainer struct {
	ContentType           *string                 `json:"content_type,omitempty"`
	CopyrightCreationTime *core.Time              `json:"copyright_creation_time,omitempty"`
	DownloadHdURL         *string                 `json:"download_hd_url,omitempty"`
	DurationInSec         *float64                `json:"duration_in_sec,omitempty"`
	ID                    *core.ID                `json:"id,omitempty"`
	Iswc                  *string                 `json:"iswc,omitempty"`
	Metadata              *map[string]interface{} `json:"metadata,omitempty"`
	PlayableVideoURI      *string                 `json:"playable_video_uri,omitempty"`
	PublishedTime         *core.Time              `json:"published_time,omitempty"`
	ThumbnailURL          *string                 `json:"thumbnail_url,omitempty"`
	Title                 *string                 `json:"title,omitempty"`
	UniversalContentID    *core.ID                `json:"universal_content_id,omitempty"`
	WriterNames           *[]string               `json:"writer_names,omitempty"`
}

var CopyrightReferenceContainerFields = struct {
	ContentType           string
	CopyrightCreationTime string
	DownloadHdURL         string
	DurationInSec         string
	ID                    string
	Iswc                  string
	Metadata              string
	PlayableVideoURI      string
	PublishedTime         string
	ThumbnailURL          string
	Title                 string
	UniversalContentID    string
	WriterNames           string
}{
	ContentType:           "content_type",
	CopyrightCreationTime: "copyright_creation_time",
	DownloadHdURL:         "download_hd_url",
	DurationInSec:         "duration_in_sec",
	ID:                    "id",
	Iswc:                  "iswc",
	Metadata:              "metadata",
	PlayableVideoURI:      "playable_video_uri",
	PublishedTime:         "published_time",
	ThumbnailURL:          "thumbnail_url",
	Title:                 "title",
	UniversalContentID:    "universal_content_id",
	WriterNames:           "writer_names",
}
