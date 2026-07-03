package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type LeadGenThankYouPageGatedFile struct {
	FileCdnURL    *string  `json:"file_cdn_url,omitempty"`
	FileName      *string  `json:"file_name,omitempty"`
	FileSizeBytes *int     `json:"file_size_bytes,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
}
