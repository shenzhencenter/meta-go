package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ImageCopyright struct {
	Artist                      *string                `json:"artist,omitempty"`
	CopyrightMonitoringStatus   *string                `json:"copyright_monitoring_status,omitempty"`
	CreationTime                *core.Time             `json:"creation_time,omitempty"`
	Creator                     *string                `json:"creator,omitempty"`
	CustomID                    *core.ID               `json:"custom_id,omitempty"`
	Description                 *string                `json:"description,omitempty"`
	Filename                    *string                `json:"filename,omitempty"`
	ID                          *core.ID               `json:"id,omitempty"`
	Image                       *Photo                 `json:"image,omitempty"`
	MatchesCount                *uint64                `json:"matches_count,omitempty"`
	OriginalContentCreationDate *core.Time             `json:"original_content_creation_date,omitempty"`
	OwnershipCountries          *VideoCopyrightGeoGate `json:"ownership_countries,omitempty"`
	Tags                        *[]string              `json:"tags,omitempty"`
	Title                       *string                `json:"title,omitempty"`
	UpdateTime                  *core.Time             `json:"update_time,omitempty"`
}

var ImageCopyrightFields = struct {
	Artist                      string
	CopyrightMonitoringStatus   string
	CreationTime                string
	Creator                     string
	CustomID                    string
	Description                 string
	Filename                    string
	ID                          string
	Image                       string
	MatchesCount                string
	OriginalContentCreationDate string
	OwnershipCountries          string
	Tags                        string
	Title                       string
	UpdateTime                  string
}{
	Artist:                      "artist",
	CopyrightMonitoringStatus:   "copyright_monitoring_status",
	CreationTime:                "creation_time",
	Creator:                     "creator",
	CustomID:                    "custom_id",
	Description:                 "description",
	Filename:                    "filename",
	ID:                          "id",
	Image:                       "image",
	MatchesCount:                "matches_count",
	OriginalContentCreationDate: "original_content_creation_date",
	OwnershipCountries:          "ownership_countries",
	Tags:                        "tags",
	Title:                       "title",
	UpdateTime:                  "update_time",
}
