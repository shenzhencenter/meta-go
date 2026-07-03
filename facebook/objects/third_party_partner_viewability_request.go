package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type ThirdPartyPartnerViewabilityRequest struct {
	CreatedTime     *time.Time                                         `json:"created_time,omitempty"`
	Description     *string                                            `json:"description,omitempty"`
	Ds              *string                                            `json:"ds,omitempty"`
	Hour            *time.Time                                         `json:"hour,omitempty"`
	ID              *core.ID                                           `json:"id,omitempty"`
	Metric          *enums.ThirdpartypartnerviewabilityrequestMetric   `json:"metric,omitempty"`
	ModifiedTime    *time.Time                                         `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                           `json:"owner_instance_id,omitempty"`
	Platform        *enums.ThirdpartypartnerviewabilityrequestPlatform `json:"platform,omitempty"`
	Status          *enums.ThirdpartypartnerviewabilityrequestStatus   `json:"status,omitempty"`
	TotalFileCount  *uint64                                            `json:"total_file_count,omitempty"`
}
