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
