package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type ImageCopyright struct {
	Artist                      *string                `json:"artist,omitempty"`
	CopyrightMonitoringStatus   *string                `json:"copyright_monitoring_status,omitempty"`
	CreationTime                *time.Time             `json:"creation_time,omitempty"`
	Creator                     *string                `json:"creator,omitempty"`
	CustomID                    *core.ID               `json:"custom_id,omitempty"`
	Description                 *string                `json:"description,omitempty"`
	Filename                    *string                `json:"filename,omitempty"`
	ID                          *core.ID               `json:"id,omitempty"`
	Image                       *Photo                 `json:"image,omitempty"`
	MatchesCount                *uint64                `json:"matches_count,omitempty"`
	OriginalContentCreationDate *time.Time             `json:"original_content_creation_date,omitempty"`
	OwnershipCountries          *VideoCopyrightGeoGate `json:"ownership_countries,omitempty"`
	Tags                        *[]string              `json:"tags,omitempty"`
	Title                       *string                `json:"title,omitempty"`
	UpdateTime                  *time.Time             `json:"update_time,omitempty"`
}
