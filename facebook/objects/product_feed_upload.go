package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"time"
)

type ProductFeedUpload struct {
	EndTime           *time.Time                          `json:"end_time,omitempty"`
	ErrorCount        *int                                `json:"error_count,omitempty"`
	ErrorReport       *ProductFeedUploadErrorReport       `json:"error_report,omitempty"`
	Filename          *string                             `json:"filename,omitempty"`
	ID                *core.ID                            `json:"id,omitempty"`
	InputMethod       *enums.ProductfeeduploadInputMethod `json:"input_method,omitempty"`
	NumDeletedItems   *int                                `json:"num_deleted_items,omitempty"`
	NumDetectedItems  *int                                `json:"num_detected_items,omitempty"`
	NumInvalidItems   *int                                `json:"num_invalid_items,omitempty"`
	NumPersistedItems *int                                `json:"num_persisted_items,omitempty"`
	StartTime         *time.Time                          `json:"start_time,omitempty"`
	URL               *string                             `json:"url,omitempty"`
	WarningCount      *int                                `json:"warning_count,omitempty"`
}
