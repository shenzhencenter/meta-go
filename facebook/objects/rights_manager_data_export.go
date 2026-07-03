package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type RightsManagerDataExport struct {
	DownloadURI    *string    `json:"download_uri,omitempty"`
	ExportScope    *string    `json:"export_scope,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	RecordType     *string    `json:"record_type,omitempty"`
	TimeRangeEnd   *time.Time `json:"time_range_end,omitempty"`
	TimeRangeStart *time.Time `json:"time_range_start,omitempty"`
}
