package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ThirdPartyPartnerViewabilityRequest struct {
	CreatedTime     *core.Time                                         `json:"created_time,omitempty"`
	Description     *string                                            `json:"description,omitempty"`
	Ds              *string                                            `json:"ds,omitempty"`
	Hour            *core.Time                                         `json:"hour,omitempty"`
	ID              *core.ID                                           `json:"id,omitempty"`
	Metric          *enums.ThirdpartypartnerviewabilityrequestMetric   `json:"metric,omitempty"`
	ModifiedTime    *core.Time                                         `json:"modified_time,omitempty"`
	OwnerInstanceID *core.ID                                           `json:"owner_instance_id,omitempty"`
	Platform        *enums.ThirdpartypartnerviewabilityrequestPlatform `json:"platform,omitempty"`
	Status          *enums.ThirdpartypartnerviewabilityrequestStatus   `json:"status,omitempty"`
	TotalFileCount  *uint64                                            `json:"total_file_count,omitempty"`
}

var ThirdPartyPartnerViewabilityRequestFields = struct {
	CreatedTime     string
	Description     string
	Ds              string
	Hour            string
	ID              string
	Metric          string
	ModifiedTime    string
	OwnerInstanceID string
	Platform        string
	Status          string
	TotalFileCount  string
}{
	CreatedTime:     "created_time",
	Description:     "description",
	Ds:              "ds",
	Hour:            "hour",
	ID:              "id",
	Metric:          "metric",
	ModifiedTime:    "modified_time",
	OwnerInstanceID: "owner_instance_id",
	Platform:        "platform",
	Status:          "status",
	TotalFileCount:  "total_file_count",
}
