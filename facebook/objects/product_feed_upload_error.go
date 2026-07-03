package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeedUploadError struct {
	AffectedSurfaces *[]enums.ProductfeeduploaderrorAffectedSurfaces `json:"affected_surfaces,omitempty"`
	Description      *string                                         `json:"description,omitempty"`
	ErrorType        *string                                         `json:"error_type,omitempty"`
	ID               *core.ID                                        `json:"id,omitempty"`
	Severity         *enums.ProductfeeduploaderrorSeverity           `json:"severity,omitempty"`
	Summary          *string                                         `json:"summary,omitempty"`
	TotalCount       *uint64                                         `json:"total_count,omitempty"`
}

var ProductFeedUploadErrorFields = struct {
	AffectedSurfaces string
	Description      string
	ErrorType        string
	ID               string
	Severity         string
	Summary          string
	TotalCount       string
}{
	AffectedSurfaces: "affected_surfaces",
	Description:      "description",
	ErrorType:        "error_type",
	ID:               "id",
	Severity:         "severity",
	Summary:          "summary",
	TotalCount:       "total_count",
}
