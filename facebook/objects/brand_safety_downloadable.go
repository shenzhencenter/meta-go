package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
