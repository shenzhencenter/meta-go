package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PartnerCenterExportFile struct {
	ID       *core.ID `json:"id,omitempty"`
	ReportDs *string  `json:"report_ds,omitempty"`
	URL      *string  `json:"url,omitempty"`
}
