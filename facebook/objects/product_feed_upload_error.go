package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
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
