package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeedUploadGet struct {
	EndTime           *interface{}                           `json:"end_time,omitempty"`
	ErrorCount        *int                                   `json:"error_count,omitempty"`
	ErrorReport       *map[string]interface{}                `json:"error_report,omitempty"`
	Errors            *map[string]interface{}                `json:"errors,omitempty"`
	Filename          *string                                `json:"filename,omitempty"`
	ID                *core.ID                               `json:"id,omitempty"`
	InputMethod       *enums.ProductfeeduploadgetInputMethod `json:"input_method,omitempty"`
	NumDeletedItems   *int                                   `json:"num_deleted_items,omitempty"`
	NumDetectedItems  *int                                   `json:"num_detected_items,omitempty"`
	NumInvalidItems   *int                                   `json:"num_invalid_items,omitempty"`
	NumPersistedItems *int                                   `json:"num_persisted_items,omitempty"`
	Progresses        *map[string]interface{}                `json:"progresses,omitempty"`
	StartTime         *interface{}                           `json:"start_time,omitempty"`
	UploadComplete    *bool                                  `json:"upload_complete,omitempty"`
	URL               *string                                `json:"url,omitempty"`
	WarningCount      *int                                   `json:"warning_count,omitempty"`
}

var ProductFeedUploadGetFields = struct {
	EndTime           string
	ErrorCount        string
	ErrorReport       string
	Errors            string
	Filename          string
	ID                string
	InputMethod       string
	NumDeletedItems   string
	NumDetectedItems  string
	NumInvalidItems   string
	NumPersistedItems string
	Progresses        string
	StartTime         string
	UploadComplete    string
	URL               string
	WarningCount      string
}{
	EndTime:           "end_time",
	ErrorCount:        "error_count",
	ErrorReport:       "error_report",
	Errors:            "errors",
	Filename:          "filename",
	ID:                "id",
	InputMethod:       "input_method",
	NumDeletedItems:   "num_deleted_items",
	NumDetectedItems:  "num_detected_items",
	NumInvalidItems:   "num_invalid_items",
	NumPersistedItems: "num_persisted_items",
	Progresses:        "progresses",
	StartTime:         "start_time",
	UploadComplete:    "upload_complete",
	URL:               "url",
	WarningCount:      "warning_count",
}
