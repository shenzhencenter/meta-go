package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeedUpload struct {
	EndTime           *core.Time                          `json:"end_time,omitempty"`
	ErrorCount        *int                                `json:"error_count,omitempty"`
	ErrorReport       *ProductFeedUploadErrorReport       `json:"error_report,omitempty"`
	Filename          *string                             `json:"filename,omitempty"`
	ID                *core.ID                            `json:"id,omitempty"`
	InputMethod       *enums.ProductfeeduploadInputMethod `json:"input_method,omitempty"`
	NumDeletedItems   *int                                `json:"num_deleted_items,omitempty"`
	NumDetectedItems  *int                                `json:"num_detected_items,omitempty"`
	NumInvalidItems   *int                                `json:"num_invalid_items,omitempty"`
	NumPersistedItems *int                                `json:"num_persisted_items,omitempty"`
	StartTime         *core.Time                          `json:"start_time,omitempty"`
	URL               *string                             `json:"url,omitempty"`
	WarningCount      *int                                `json:"warning_count,omitempty"`
}

var ProductFeedUploadFields = struct {
	EndTime           string
	ErrorCount        string
	ErrorReport       string
	Filename          string
	ID                string
	InputMethod       string
	NumDeletedItems   string
	NumDetectedItems  string
	NumInvalidItems   string
	NumPersistedItems string
	StartTime         string
	URL               string
	WarningCount      string
}{
	EndTime:           "end_time",
	ErrorCount:        "error_count",
	ErrorReport:       "error_report",
	Filename:          "filename",
	ID:                "id",
	InputMethod:       "input_method",
	NumDeletedItems:   "num_deleted_items",
	NumDetectedItems:  "num_detected_items",
	NumInvalidItems:   "num_invalid_items",
	NumPersistedItems: "num_persisted_items",
	StartTime:         "start_time",
	URL:               "url",
	WarningCount:      "warning_count",
}
