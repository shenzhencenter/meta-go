package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PartnerCenterExportFile struct {
	ID       *core.ID `json:"id,omitempty"`
	ReportDs *string  `json:"report_ds,omitempty"`
	URL      *string  `json:"url,omitempty"`
}

var PartnerCenterExportFileFields = struct {
	ID       string
	ReportDs string
	URL      string
}{
	ID:       "id",
	ReportDs: "report_ds",
	URL:      "url",
}
