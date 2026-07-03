package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ThirdPartyMeasurementReportDataset struct {
	Category *string                   `json:"category,omitempty"`
	ID       *core.ID                  `json:"id,omitempty"`
	Partner  *Business                 `json:"partner,omitempty"`
	Product  *string                   `json:"product,omitempty"`
	Schema   *[]map[string]interface{} `json:"schema,omitempty"`
}
