package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BrandSafetyDownloadable struct {
	AccountContextID        *core.ID `json:"account_context_id,omitempty"`
	AsyncJobPercentComplete *int     `json:"async_job_percent_complete,omitempty"`
	AsyncJobStatus          *string  `json:"async_job_status,omitempty"`
	FileName                *string  `json:"file_name,omitempty"`
	ID                      *core.ID `json:"id,omitempty"`
	RequestSurface          *string  `json:"request_surface,omitempty"`
	URL                     *string  `json:"url,omitempty"`
}

var BrandSafetyDownloadableFields = struct {
	AccountContextID        string
	AsyncJobPercentComplete string
	AsyncJobStatus          string
	FileName                string
	ID                      string
	RequestSurface          string
	URL                     string
}{
	AccountContextID:        "account_context_id",
	AsyncJobPercentComplete: "async_job_percent_complete",
	AsyncJobStatus:          "async_job_status",
	FileName:                "file_name",
	ID:                      "id",
	RequestSurface:          "request_surface",
	URL:                     "url",
}
