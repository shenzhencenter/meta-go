package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type OrderIDAttributions struct {
	AppID            *core.ID                  `json:"app_id,omitempty"`
	AttributionType  *string                   `json:"attribution_type,omitempty"`
	Attributions     *[]map[string]interface{} `json:"attributions,omitempty"`
	ConversionDevice *string                   `json:"conversion_device,omitempty"`
	DatasetID        *core.ID                  `json:"dataset_id,omitempty"`
	HoldoutStatus    *[]map[string]interface{} `json:"holdout_status,omitempty"`
	OrderID          *core.ID                  `json:"order_id,omitempty"`
	OrderTimestamp   *time.Time                `json:"order_timestamp,omitempty"`
	PixelID          *core.ID                  `json:"pixel_id,omitempty"`
}
