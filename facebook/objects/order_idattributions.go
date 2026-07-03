package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OrderIDAttributions struct {
	AppID            *core.ID                  `json:"app_id,omitempty"`
	AttributionType  *string                   `json:"attribution_type,omitempty"`
	Attributions     *[]map[string]interface{} `json:"attributions,omitempty"`
	ConversionDevice *string                   `json:"conversion_device,omitempty"`
	DatasetID        *core.ID                  `json:"dataset_id,omitempty"`
	HoldoutStatus    *[]map[string]interface{} `json:"holdout_status,omitempty"`
	OrderID          *core.ID                  `json:"order_id,omitempty"`
	OrderTimestamp   *core.Time                `json:"order_timestamp,omitempty"`
	PixelID          *core.ID                  `json:"pixel_id,omitempty"`
}

var OrderIDAttributionsFields = struct {
	AppID            string
	AttributionType  string
	Attributions     string
	ConversionDevice string
	DatasetID        string
	HoldoutStatus    string
	OrderID          string
	OrderTimestamp   string
	PixelID          string
}{
	AppID:            "app_id",
	AttributionType:  "attribution_type",
	Attributions:     "attributions",
	ConversionDevice: "conversion_device",
	DatasetID:        "dataset_id",
	HoldoutStatus:    "holdout_status",
	OrderID:          "order_id",
	OrderTimestamp:   "order_timestamp",
	PixelID:          "pixel_id",
}
